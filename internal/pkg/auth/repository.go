package auth

import (
	"context"

	authv1 "github.com/MyyPo/w34.Go/gen/go/auth/v1"
	"github.com/MyyPo/w34.Go/gen/psql/auth/public/model"
)

type Repository interface {
	CreateUser(ctx context.Context, req *authv1.SignUpRequest) (model.Accounts, error)
}
