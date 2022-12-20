package dev

import (
	"context"
	"fmt"

	devv1 "github.com/MyyPo/w34.Go/gen/go/dev/v1"
	"github.com/MyyPo/w34.Go/gen/psql/main/public/model"
	validator "github.com/MyyPo/w34.Go/internal/pkg/dev/validator"
	"google.golang.org/grpc/metadata"
)

type DevServer struct {
	repo      Repository
	validator validator.DevValidator
}

func NewDevServer(repo Repository, validator validator.DevValidator) *DevServer {
	return &DevServer{
		repo:      repo,
		validator: validator,
	}
}

func (s DevServer) CreateProject(
	ctx context.Context,
	req *devv1.NewProjectRequest,
) (*devv1.NewProjectResponse, error) {
	reqUserID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	projectName := req.GetName()

	if err = s.validator.ValidateName(projectName); err != nil {
		return nil, err
	}

	_, err = s.repo.CreateProject(ctx, projectName, reqUserID)
	if err != nil {
		return nil, err
	}

	return &devv1.NewProjectResponse{}, nil
}

func (s DevServer) DeleteProject(
	ctx context.Context,
	req *devv1.DeleteProjectRequest,
) (*devv1.DeleteProjectResponse, error) {
	reqUserID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	projectName := req.GetName()

	err = s.repo.DeleteProject(ctx, projectName, reqUserID)
	if err != nil {
		return nil, err
	}

	return &devv1.DeleteProjectResponse{}, nil
}

func (s DevServer) CreateLocation(
	ctx context.Context,
	req *devv1.NewLocationRequest,
) (*devv1.NewLocationResponse, error) {
	reqUserID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	projectName := req.GetProjectName()
	locationName := req.GetLocationName()

	if err = s.validator.ValidateName(locationName); err != nil {
		return nil, err
	}

	_, err = s.repo.CreateLocation(ctx, projectName, locationName, reqUserID)
	if err != nil {
		return nil, err
	}

	return &devv1.NewLocationResponse{}, nil
}

func (s DevServer) CreateScene(
	ctx context.Context,
	req *devv1.NewSceneRequest,
) (*devv1.NewSceneResponse, error) {
	reqUserID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	projectName := req.GetProject()
	locationName := req.GetLocation()
	sceneOptions := req.GetOptions()

	res, err := s.repo.CreateScene(ctx, projectName, locationName, reqUserID, sceneOptions)
	if err != nil {
		return nil, err
	}

	return &devv1.NewSceneResponse{
		SceneId: res.ID,
	}, nil
}

func (s DevServer) GetLocationScenes(
	ctx context.Context,
	req *devv1.GetLocationScenesRequest,
) (*devv1.GetLocationScenesResponse, error) {
	reqUserID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	projectName := req.GetProject()
	locationName := req.GetLocation()

	modelScenes, err := s.repo.GetLocationScenes(ctx, projectName, locationName, reqUserID)
	if err != nil {
		return nil, err
	}
	res := getScenesFromModels(modelScenes)

	return &devv1.GetLocationScenesResponse{
		Scene: res,
	}, nil
}

func getScenesFromModels(modelScenes []model.Scenes) []*devv1.Scene {
	var scenes []*devv1.Scene
	for _, v := range modelScenes {
		scene := devv1.Scene{
			Options: v.Options,
		}
		scenes = append(scenes, &scene)
	}
	return scenes
}

func getUserID(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("metadata was not provided")
	}
	userIDArr := md["user_id"]
	if len(userIDArr) == 0 {
		return "", fmt.Errorf("metadata was not provided")
	}
	userID := userIDArr[0]

	return userID, nil
}
