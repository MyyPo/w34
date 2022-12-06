package auth

import (
	"context"

	authv1 "github.com/MyyPo/w34.Go/gen/go/auth/v1"
	"github.com/MyyPo/w34.Go/gen/psql/auth/public/model"
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

func (s AuthServer) SignUp(ctx context.Context, req *authv1.SignUpRequest) (model.Accounts, error) {
	err := s.validator.ValidateCredentials(req)
	if err != nil {
		return model.Accounts{}, err
	}

	return s.repo.CreateUser(ctx, req)
}
