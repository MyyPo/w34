package dev

import (
	"context"
	"database/sql"
	"fmt"

	"log"
	"testing"

	"github.com/MyyPo/w34.Go/configs"
	devv1 "github.com/MyyPo/w34.Go/gen/go/dev/v1"
	t "github.com/MyyPo/w34.Go/gen/psql/main/public/table"
	"github.com/MyyPo/w34.Go/internal/adapters/auth/psql"
	"github.com/MyyPo/w34.Go/internal/adapters/dev/psql"
	"github.com/MyyPo/w34.Go/internal/jwt"
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
	devServer, accessToken := setupPsql(t)

	t.Run("Valid create a new project", func(t *testing.T) {
		md := metadata.MD{
			"access_token": []string{accessToken},
		}
		ctx := metadata.NewIncomingContext(context.Background(), md)

		req := &devv1.NewProjectRequest{
			Name:   "int_test",
			Public: true,
		}

		_, err := devServer.CreateProject(ctx, req)
		if err != nil {
			t.Errorf("error creating project: %v", err)
		}

	})
	t.Run("Try to create a new project with the same name", func(t *testing.T) {
		md := metadata.MD{
			"access_token": []string{accessToken},
		}
		ctx := metadata.NewIncomingContext(context.Background(), md)

		req := &devv1.NewProjectRequest{
			Name:   "int_test",
			Public: true,
		}

		_, err := devServer.CreateProject(ctx, req)
		if err == nil {
			t.Errorf("created a project with the repeating name")
		}

	})
	t.Run("Delete the created project", func(t *testing.T) {
		md := metadata.MD{
			"access_token": []string{accessToken},
		}
		ctx := metadata.NewIncomingContext(context.Background(), md)

		req := &devv1.DeleteProjectRequest{
			Name: "int_test",
		}

		_, err := devServer.DeleteProject(ctx, req)
		if err != nil {
			t.Errorf("failed to delete a project: %v", err)
		}
	})
	t.Run("Try to delete the deleted project again", func(t *testing.T) {
		md := metadata.MD{
			"access_token": []string{accessToken},
		}
		ctx := metadata.NewIncomingContext(context.Background(), md)

		req := &devv1.DeleteProjectRequest{
			Name: "int_test",
		}

		_, err := devServer.DeleteProject(ctx, req)
		if err == nil {
			t.Errorf("no error raised trying to delete the project")
		}
	})

}

func removeRows(db *sql.DB, testUserID int32) {
	id64 := int64(testUserID)

	stmt := t.Projects.
		DELETE().
		WHERE(
			t.Projects.OwnerID.EQ(Int(id64)),
		)
	stmt.Exec(db)

	stmt = t.Accounts.
		DELETE().
		WHERE(
			t.Accounts.UserID.EQ(Int(id64)),
		)
	stmt.Exec(db)
}

func setupPsql(t *testing.T) (*DevServer, string) {

	config, err := configs.NewConfig("../../../configs")
	if err != nil {
		t.Errorf("failed to load config: %q", err)
	}

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

	testJWTManager, err := jwt.NewJWTManager("../../../configs/rsa", "../../../configs/rsa.pub",
		"../../../configs/refresh_rsa", "../../../configs/refresh_rsa.pub",
		config.AccessTokenDuration, config.RefreshTokenDuration)
	if err != nil {
		log.Fatalf("failed to create jwtManager for testing: %q", err)
	}

	testAccessToken, err := testJWTManager.GenerateAccessToken(testUser.UserID)
	if err != nil {
		log.Fatalf("failed to create accesstoken for testing: %q", err)
	}

	psqlRepo := dev_psql_adapter.NewDevPSQLRepository(psqlDB)

	// remove all affected database rows after the tests
	t.Cleanup(func() { removeRows(psqlDB, testUser.UserID) })
	return NewDevServer(psqlRepo, *testJWTManager), testAccessToken
}
