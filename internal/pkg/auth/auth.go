package auth

import (
	"context"

	authv1 "github.com/MyyPo/w34.Go/gen/go/auth/v1"
	"github.com/MyyPo/w34.Go/internal/pkg/auth/hasher"
	"github.com/MyyPo/w34.Go/internal/pkg/auth/redis"
	"github.com/MyyPo/w34.Go/internal/pkg/auth/validators"
)

type AuthServer struct {
	repo        Repository
	redisClient auth_redis.RedisClient
	validator   validators.AuthValidator
	jwtManager  JWTManager
	hasher      hasher.Hasher
	us          authv1.UnimplementedAuthServiceServer
}

func NewAuthServer(
	repo Repository,
	redisClient auth_redis.RedisClient,
	validator validators.AuthValidator,
	jwtManager JWTManager,
) *AuthServer {
	return &AuthServer{
		repo:        repo,
		redisClient: redisClient,
		validator:   validator,
		jwtManager:  jwtManager,
		us:          authv1.UnimplementedAuthServiceServer{},
	}
}

func (s AuthServer) SignUp(ctx context.Context, req *authv1.SignUpRequest) (*authv1.SignUpResponse, error) {
	newUsername := req.GetUsername()
	newEmail := req.GetEmail()
	newPassword := req.GetPassword()

	if err := s.validator.ValidateSignUpCredentials(
		newUsername,
		newEmail,
		newPassword,
	); err != nil {
		return nil, err
	}

	newHashedPassword, err := s.hasher.HashSecret(newPassword)
	if err != nil {
		return nil, err
	}

	createdUser, err := s.repo.CreateUser(ctx, newUsername, newEmail, newHashedPassword)
	if err != nil {
		return nil, err
	}

	createdUserID := createdUser.UserID

	accessToken, err := s.jwtManager.GenerateAccessToken(createdUserID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtManager.GenerateRefreshToken(createdUserID)
	if err != nil {
		return nil, err
	}

	err = s.redisClient.StoreRefreshToken(ctx, createdUserID, refreshToken)
	if err != nil {
		return nil, err
	}

	return &authv1.SignUpResponse{
		Tokens: &authv1.TokenPackage{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func (s AuthServer) SignIn(ctx context.Context, req *authv1.SignInRequest) (*authv1.SignInResponse, error) {
	usernameOrEmail := req.GetUnOrEmail()
	userIDAndPassword, err := s.repo.LookupExistingUser(ctx, usernameOrEmail)
	// throw error if user not found
	if err != nil {
		return nil, err
	}
	reqPassword := req.GetPassword()
	hashedUserPassword := userIDAndPassword.Password

	// compared password from request with hash returned from database
	if err := s.hasher.CompareWithSecret(reqPassword, hashedUserPassword); err != nil {
		return nil, err
	}

	retrievedUserID := userIDAndPassword.UserID

	accessToken, err := s.jwtManager.GenerateAccessToken(retrievedUserID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtManager.GenerateRefreshToken(retrievedUserID)
	if err != nil {
		return nil, err
	}

	err = s.redisClient.StoreRefreshToken(ctx, retrievedUserID, refreshToken)
	if err != nil {
		return nil, err
	}

	return &authv1.SignInResponse{
		Tokens: &authv1.TokenPackage{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
