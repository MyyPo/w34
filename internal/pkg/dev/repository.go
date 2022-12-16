package dev

import (
	"context"

	"github.com/MyyPo/w34.Go/gen/psql/main/public/model"
)

type Repository interface {
	CreateProject(
		ctx context.Context,
		projectName string,
		ownerID string,
	) (model.Projects, error)
	DeleteProject(
		ctx context.Context,
		projectName string,
		ownerID string,
	) error
	CreateLocation(
		ctx context.Context,
		projectID int32,
		locationName string,
	) (model.Locations, error)
}
