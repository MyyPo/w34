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
		reqUserID string,
	) error
	CreateLocation(
		ctx context.Context,
		projectName string,
		locationName string,
		reqUserID string,
	) (model.Locations, error)
	CreateScene(
		ctx context.Context,
		projectName string,
		locationName string,
		reqUserID string,
		sceneOptions string,
	) (model.Scenes, error)
}