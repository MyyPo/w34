package auth

import (
	"context"

	"github.com/MyyPo/w34.Go/gen/psql/main/public/model"
)

type Repository interface {
	CreateUser(
		ctx context.Context,
		newUsername string,
		newEmail string,
		newHashedPassword string,
	) (model.Accounts, error)
	LookupExistingUser(
		ctx context.Context,
		usernameOrEmail string,
	) (model.Accounts, error)
}
