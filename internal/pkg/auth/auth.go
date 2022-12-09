package auth

import (
	"context"

	authv1 "github.com/MyyPo/w34.Go/gen/go/auth/v1"
	"github.com/MyyPo/w34.Go/internal/pkg/validators"
)

type AuthServer struct {
	repo       Repository
	validator  validators.AuthValidator
	jwtManager JWTManager
	hasher     Hasher
	us         authv1.UnimplementedAuthServiceServer
}

func NewAuthServer(repo Repository, validator validators.AuthValidator, jwtManager JWTManager) *AuthServer {
	return &AuthServer{
		repo:       repo,
		validator:  validator,
		jwtManager: jwtManager,
		us:         authv1.UnimplementedAuthServiceServer{},
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

	accessToken, err := s.jwtManager.GenerateAccessToken(createdUser.UserID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtManager.GenerateRefreshToken(createdUser.UserID)
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
	return nil, nil
}
