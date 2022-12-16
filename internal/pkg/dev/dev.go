package dev

import (
	"context"
	"fmt"

	devv1 "github.com/MyyPo/w34.Go/gen/go/dev/v1"
	"github.com/MyyPo/w34.Go/internal/jwt"
	"google.golang.org/grpc/metadata"
)

type DevServer struct {
	repo       Repository
	jwtManager jwt.JWTManager
}

func NewDevServer(repo Repository, jwtManager jwt.JWTManager) *DevServer {
	return &DevServer{
		repo:       repo,
		jwtManager: jwtManager,
	}
}

func (s DevServer) CreateProject(
	ctx context.Context,
	req *devv1.NewProjectRequest,
) (*devv1.NewProjectResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata was not provided")
	}
	accessArr := md["access_token"]
	if len(accessArr) == 0 {
		return nil, fmt.Errorf("access token was not provided")
	}
	accessToken := accessArr[0]

	tokenClaims, err := s.jwtManager.ValidateJwtExtractClaims(accessToken, s.jwtManager.AccessPublicSignature)
	if err != nil {
		return nil, err
	}
	ownerID := tokenClaims.Subject

	projectName := req.GetName()

	_, err = s.repo.CreateProject(ctx, projectName, ownerID)
	if err != nil {
		return nil, err
	}

	return &devv1.NewProjectResponse{}, nil
}

func (s DevServer) DeleteProject(
	ctx context.Context,
	req *devv1.DeleteProjectRequest,
) (*devv1.DeleteProjectResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata was not provided")
	}
	accessArr := md["access_token"]
	if len(accessArr) == 0 {
		return nil, fmt.Errorf("access token was not provided")
	}
	accessToken := accessArr[0]

	tokenClaims, err := s.jwtManager.ValidateJwtExtractClaims(accessToken, s.jwtManager.AccessPublicSignature)
	if err != nil {
		return nil, err
	}
	ownerID := tokenClaims.Subject

	projectName := req.GetName()

	err = s.repo.DeleteProject(ctx, projectName, ownerID)
	if err != nil {
		return nil, err
	}

	return &devv1.DeleteProjectResponse{}, nil
}

func (s DevServer) CreateLocation(
	ctx context.Context,
	req *devv1.NewLocationRequest,
) (*devv1.NewLocationResponse, error) {
	return nil, nil
}
