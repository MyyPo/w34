package dev

import "context"

type Repository interface {
	CreateProject(
		ctx context.Context,
		projectName string,
	)
}
