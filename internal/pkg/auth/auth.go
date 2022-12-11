package auth

import (
	"context"
	"fmt"
	"strconv"

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

	accessToken, err := s.jwtManager.GenerateAccessToken(createdUserID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtManager.GenerateRefreshToken(createdUserID)
	if err != nil {
		return nil, err
	}

	// hashedRefreshToken, err := s.hasher.HashSecret(refreshToken)
	// if err != nil {
	// 	return nil, err
	// }

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

	// delete the valid refresh token stored in db for this account, if it exists
	err = s.redisClient.DeleteRefreshToken(ctx, retrievedUserID)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.jwtManager.GenerateAccessToken(retrievedUserID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtManager.GenerateRefreshToken(retrievedUserID)
	if err != nil {
		return nil, err
	}

	// hashedRefreshToken, err := s.hasher.HashSecret(refreshToken)
	// if err != nil {
	// 	return nil, err
	// }

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

func (s AuthServer) RefreshTokens(
	ctx context.Context,
	req *authv1.RefreshTokensRequest,
) (*authv1.RefreshTokensResponse, error) {

	reqRefreshToken := req.GetRefreshToken()
	tokenClaims, err := s.jwtManager.ValidateJwtExtractClaims(reqRefreshToken, s.jwtManager.pathToRefreshPublicSignature)
	if err != nil {
		return nil, err
	}

	userID := tokenClaims.Subject

	// lookup if this refresh token was really in the database
	currentTokenInDB, err := s.redisClient.GetToken(ctx, userID)
	if err != nil {
		return nil, err
	}
	// throw an error if the token isn't stored in redis
	if currentTokenInDB != reqRefreshToken {
		return nil, fmt.Errorf("used refresh token was provided")
	}

	// if err := s.hasher.CompareWithSecret(currentTokenInDB, reqRefreshToken); err != nil {
	// 	return nil, fmt.Errorf("used refresh token was provided")
	// }

	// delete the current refresh token stored in db for this account
	err = s.redisClient.DeleteRefreshTokenStringID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// create new tokens
	intUserID, _ := strconv.ParseInt(userID, 10, 32)
	int32UserID := int32(intUserID)
	newAccessToken, err := s.jwtManager.GenerateAccessToken(int32UserID)
	if err != nil {
		return nil, err
	}
	newRefreshToken, err := s.jwtManager.GenerateRefreshToken(int32UserID)
	if err != nil {
		return nil, err
	}

	// newHashedRefreshToken, err := s.hasher.HashSecret(newRefreshToken)
	// if err != nil {
	// 	return nil, err
	// }

	err = s.redisClient.StoreRefreshTokenStringID(ctx, userID, newRefreshToken)
	if err != nil {
		return nil, err
	}

	return &authv1.RefreshTokensResponse{
		Tokens: &authv1.TokenPackage{
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken,
		},
	}, nil
}
