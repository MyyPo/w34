package auth

import (
	"context"

	authv1 "github.com/MyyPo/w34.Go/gen/go/auth/v1"
	"github.com/MyyPo/w34.Go/gen/psql/auth/public/model"
)

type AuthServer struct {
	repo Repository
	us   authv1.UnimplementedAuthServiceServer
}

func NewAuthServer(repo Repository) *AuthServer {
	return &AuthServer{
		repo: repo,
		us:   authv1.UnimplementedAuthServiceServer{},
	}
}

func (s AuthServer) SignUp(ctx context.Context, req *authv1.SignUpRequest) (model.Accounts, error) {
	return s.repo.CreateUser(ctx, req)
}
