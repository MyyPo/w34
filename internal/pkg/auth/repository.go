package auth

import (
	"context"

	"github.com/MyyPo/w34.Go/gen/psql/auth/public/model"
)

type Repository interface {
	CreateUser(ctx context.Context,
		newUsername string,
		newEmail string,
		newHashedPassword string,
	) (model.Accounts, error)
	// LookupUser(ctx context.Context, req *authv1.SignInRequest) (model.Accounts, error)
}
