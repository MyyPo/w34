package dev

import (
	"context"
	"fmt"

	devv1 "github.com/MyyPo/w34.Go/gen/go/dev/v1"
	"google.golang.org/grpc/metadata"
)

type DevServer struct {
	repo Repository
}

func NewDevServer(repo Repository) *DevServer {
	return &DevServer{
		repo: repo,
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

	_, err = s.repo.CreateLocation(ctx, projectName, locationName, reqUserID)
	if err != nil {
		return nil, err
	}

	return &devv1.NewLocationResponse{}, nil
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
