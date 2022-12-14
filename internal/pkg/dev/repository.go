package dev

import (
	"context"

	"github.com/MyyPo/w34.Go/gen/psql/main/public/model"
)

type Repository interface {
	CreateProject(
		ctx context.Context,
		projectName string,
	) (model.Projects, error)
}
