package auth

import (
	"context"
	"fmt"
	"strconv"

	authv1 "github.com/MyyPo/w34.Go/gen/go/auth/v1"
	"github.com/MyyPo/w34.Go/internal/pkg/auth/hasher"
	"github.com/MyyPo/w34.Go/internal/pkg/auth/jwt"
	"github.com/MyyPo/w34.Go/internal/pkg/auth/redis"
	"github.com/MyyPo/w34.Go/internal/pkg/auth/validators"
)

type AuthServer struct {
	repo        Repository
	redisClient auth_redis.RedisClient
	validator   validators.AuthValidator
	jwtManager  jwt.JWTManager
	hasher      hasher.Hasher
	us          authv1.UnimplementedAuthServiceServer
}

func NewAuthServer(
	repo Repository,
	redisClient auth_redis.RedisClient,
	validator validators.AuthValidator,
	jwtManager jwt.JWTManager,
	hasher hasher.Hasher,
) *AuthServer {
	return &AuthServer{
		repo:        repo,
		redisClient: redisClient,
		validator:   validator,
		jwtManager:  jwtManager,
		hasher:      hasher,
		us:          authv1.UnimplementedAuthServiceServer{},
	}
}

func (s AuthServer) SignUp(
	ctx context.Context,
	req *authv1.SignUpRequest,
) (*authv1.SignUpResponse, error) {
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

	newAccessToken, newRefreshToken, err := s.createTokensAndStoreRefresh(ctx, createdUserID)
	if err != nil {
		return nil, fmt.Errorf("internal error")
	}

	return &authv1.SignUpResponse{
		Tokens: &authv1.TokenPackage{
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken,
		},
	}, nil
}

func (s AuthServer) SignIn(
	ctx context.Context,
	req *authv1.SignInRequest,
) (*authv1.SignInResponse, error) {
	usernameOrEmail := req.GetUnOrEmail()
	userIDAndPassword, err := s.repo.LookupExistingUser(ctx, usernameOrEmail)
	// throw error if user not found
	if err != nil {
		return nil, err
	}
	reqPassword := req.GetPassword()
	hashedUserPassword := userIDAndPassword.Password

	// compared password from request with hash returned from database
	if err := s.hasher.CompareWithSecret(hashedUserPassword, reqPassword); err != nil {
		return nil, err
	}

	retrievedUserID := userIDAndPassword.UserID

	newAccessToken, newRefreshToken, err := s.createTokensAndStoreRefresh(ctx, retrievedUserID)
	if err != nil {
		return nil, fmt.Errorf("internal error")
	}

	return &authv1.SignInResponse{
		Tokens: &authv1.TokenPackage{
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken,
		},
	}, nil
}

func (s AuthServer) RefreshTokens(
	ctx context.Context,
	req *authv1.RefreshTokensRequest,
) (*authv1.RefreshTokensResponse, error) {

	reqRefreshToken := req.GetRefreshToken()
	tokenClaims, err := s.jwtManager.ValidateJwtExtractClaims(reqRefreshToken, s.jwtManager.RefreshPublicSignature)
	if err != nil {
		return nil, err
	}

	userID := tokenClaims.Subject

	// lookup if this refresh token was really in the database
	currentTokenInDB, err := s.redisClient.GetToken(ctx, userID)
	if err != nil {
		return nil, err
	}
	// throw an error if the token from request isn't stored in redis
	if currentTokenInDB != reqRefreshToken {
		return nil, fmt.Errorf("used refresh token was provided")
	}

	// create new tokens
	intUserID, _ := strconv.ParseInt(userID, 10, 32)
	int32UserID := int32(intUserID)
	newAccessToken, newRefreshToken, err := s.createTokensAndStoreRefresh(ctx, int32UserID)
	if err != nil {
		return nil, fmt.Errorf("internal error")
	}

	return &authv1.RefreshTokensResponse{
		Tokens: &authv1.TokenPackage{
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken,
		},
	}, nil
}

func (s AuthServer) createTokensAndStoreRefresh(
	ctx context.Context,
	userID int32,

) (
	newAccessToken string,
	newRefreshToken string,
	err error,
) {
	newAccessToken, err = s.jwtManager.GenerateAccessToken(userID)
	if err != nil {
		return "", "", err
	}
	newRefreshToken, err = s.jwtManager.GenerateRefreshToken(userID)
	if err != nil {
		return "", "", err
	}

	err = s.redisClient.StoreRefreshToken(ctx, userID, newRefreshToken)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}
