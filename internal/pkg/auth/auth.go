package auth

import (
	"context"

	authv1 "github.com/MyyPo/w34.Go/gen/go/auth/v1"
	"github.com/MyyPo/w34.Go/internal/pkg/validators"
)

type AuthServer struct {
	repo      Repository
	validator validators.AuthValidator
	us        authv1.UnimplementedAuthServiceServer
}

func NewAuthServer(repo Repository, validator validators.AuthValidator) *AuthServer {
	return &AuthServer{
		repo:      repo,
		validator: validator,
		us:        authv1.UnimplementedAuthServiceServer{},
	}
}

func (s AuthServer) SignUp(ctx context.Context, req *authv1.SignUpRequest) (*authv1.SignUpResponse, error) {
	if err := s.validator.ValidateCredentials(req); err != nil {
		return nil, err
	}
	if _, err := s.repo.CreateUser(ctx, req); err != nil {
		return nil, err
	}

	return &authv1.SignUpResponse{
		Tokens: &authv1.TokenPackage{
			AccessToken:  "",
			RefreshToken: "",
		},
	}, nil
}
