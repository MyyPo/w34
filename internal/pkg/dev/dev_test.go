package dev

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"log"
	"testing"

	devv1 "github.com/MyyPo/w34.Go/gen/go/dev/v1"
	t "github.com/MyyPo/w34.Go/gen/psql/main/public/table"
	"github.com/MyyPo/w34.Go/internal/adapters/auth/psql"
	"github.com/MyyPo/w34.Go/internal/adapters/dev/psql"
	validator "github.com/MyyPo/w34.Go/internal/pkg/dev/validator"
	. "github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/metadata"
)

const (
	host     = "host.docker.internal"
	port     = 5432
	user     = "spuser"
	password = "SPuser96"
	dbname   = "auth"
)

func TestDevServer(t *testing.T) {
	devServer, testUserID := setupPsql(t)
	strTestUserID := strconv.FormatInt(int64(testUserID), 10)
	projectName := "integr test"
	locationName := "Imperial city sewers"
	var sceneID int32

	md := metadata.MD{
		"user_id": []string{strTestUserID},
	}
	ctx := metadata.NewIncomingContext(context.Background(), md)

	t.Run("Valid create a new project", func(t *testing.T) {
		req := &devv1.NewProjectRequest{
			Name:     projectName,
			IsPublic: true,
		}

		_, err := devServer.CreateProject(ctx, req)
		if err != nil {
			t.Errorf("error creating project: %v", err)
		}
	})
	t.Run("Try to create a new project with the same name", func(t *testing.T) {
		req := &devv1.NewProjectRequest{
			Name:     projectName,
			IsPublic: true,
		}

		_, err := devServer.CreateProject(ctx, req)
		if err == nil {
			t.Errorf("created a project with the repeating name")
		}

	})
	t.Run("Create a new location", func(t *testing.T) {
		req := &devv1.NewLocationRequest{
			ProjectName:  projectName,
			LocationName: locationName,
		}

		_, err := devServer.CreateLocation(ctx, req)
		if err != nil {
			t.Errorf("failed to create a valid location: %v", err)
		}
	})
	t.Run("Try to create a new location with another user's id", func(t *testing.T) {
		md := metadata.MD{
			"user_id": []string{"1337"},
		}
		ctx := metadata.NewIncomingContext(context.Background(), md)

		req := &devv1.NewLocationRequest{
			ProjectName:  projectName,
			LocationName: "Very bad place",
		}

		_, err := devServer.CreateLocation(ctx, req)
		if err == nil {
			t.Errorf("created context with token from another user")
		}

		t.Log(err)
	})
	t.Run("Create a new scene", func(t *testing.T) {
		req := &devv1.NewSceneRequest{
			Project:  projectName,
			Location: locationName,
			Options: map[string]string{
				"A1": "NEXT 3",
				"A2": "ADD 19",
			},
		}

		res, err := devServer.CreateScene(ctx, req)
		if err != nil {
			t.Errorf("unexpected error creating a valid scene: %v", err)
		}
		sceneID = res.GetSceneId()
	})
	t.Run("Get all location scenes", func(t *testing.T) {
		req := &devv1.GetLocationScenesRequest{
			Project:  projectName,
			Location: locationName,
		}

		res, err := devServer.GetLocationScenes(ctx, req)
		if err != nil {
			t.Errorf("unexpected error retrieveing loc scenes: %v", err)
		}
		t.Logf("acquired list of scenes %v", res)
	})

	t.Run("Delete a scene", func(t *testing.T) {
		req := &devv1.DeleteSceneRequest{
			Project:  projectName,
			Location: locationName,
			SceneId:  sceneID,
		}

		_, err := devServer.DeleteScene(ctx, req)
		if err != nil {
			t.Errorf("unexpected error trying to delete a scene: %v", err)
		}
	})
	t.Run("Try to delete a scene with invalid locationName", func(t *testing.T) {
		req := &devv1.DeleteSceneRequest{
			Project:  projectName,
			Location: "wrong",
			SceneId:  sceneID,
		}

		_, err := devServer.DeleteScene(ctx, req)
		if err == nil {
			t.Errorf("expected to rasie an error passing invalid Location")
		}
	})

	t.Run("Delete the created project", func(t *testing.T) {
		req := &devv1.DeleteProjectRequest{
			Name: projectName,
		}

		_, err := devServer.DeleteProject(ctx, req)
		if err != nil {
			t.Errorf("failed to delete a project: %v", err)
		}
	})
	t.Run("Try to delete the deleted project again", func(t *testing.T) {
		req := &devv1.DeleteProjectRequest{
			Name: projectName,
		}

		_, err := devServer.DeleteProject(ctx, req)
		if err == nil {
			t.Errorf("no error raised trying to delete the project")
		}
	})

}

func removeRows(db *sql.DB, testUserID int32) {
	id64 := int64(testUserID)

	stmt := t.Accounts.
		DELETE().
		WHERE(
			t.Accounts.UserID.EQ(Int(id64)),
		)
	stmt.Exec(db)
}

func setupPsql(t *testing.T) (*DevServer, int32) {
	psqlDB, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname))
	if err != nil {
		log.Fatalf("failed to connect to db for testing: %q", err)
	}
	testAuthRep := auth_psql_adapter.NewAuthPSQLRepository(psqlDB)
	testUser, err := testAuthRep.CreateUser(context.Background(), "tester", "test@gmail.com", "testpassword")
	if err != nil {
		log.Fatalf("failed to create test user for testing: %q", err)
	}

	psqlRepo := dev_psql_adapter.NewDevPSQLRepository(psqlDB)

	devValidator, _ := validator.NewDevValidator("")

	// remove all affected database rows after the tests
	t.Cleanup(func() { removeRows(psqlDB, testUser.UserID) })
	return NewDevServer(psqlRepo, *devValidator), testUser.UserID
}
