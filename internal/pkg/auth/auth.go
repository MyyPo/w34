package auth

import (
	"context"

	authv1 "github.com/MyyPo/w34.Go/gen/go/auth/v1"
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

func (s AuthServer) SignUp(ctx context.Context) string {
	return s.repo.CreateUser(ctx)
}
