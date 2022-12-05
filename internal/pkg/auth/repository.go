package auth

import "context"

type Repository interface {
	CreateUser(ctx context.Context) string
}
