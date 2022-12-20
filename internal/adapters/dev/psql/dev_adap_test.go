package dev_psql_adapter

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	t "github.com/MyyPo/w34.Go/gen/psql/main/public/table"
	j "github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "host.docker.internal"
	port     = 5432
	user     = "spuser"
	password = "SPuser96"
	dbname   = "auth"
)

func TestDevAdapter(t *testing.T) {
	var projectName = "test"
	var locationName = "Test location"
	var projectID int32
	ownerID := "47"
	var sceneID int32

	psqlDB, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname))
	if err != nil {
		log.Fatalf("failed to connect to db for testing: %q", err)
	}

	psqlRepo := NewDevPSQLRepository(psqlDB)

	t.Run("Valid create new project", func(t *testing.T) {
		got, err := psqlRepo.CreateProject(context.Background(), projectName, ownerID)
		if err != nil {
			t.Errorf("undexpected error: %v", err)

		}

		projectID = got.ID
	})

	t.Run("Valid create a new location", func(t *testing.T) {
		got, err := psqlRepo.CreateLocation(context.Background(), projectName, locationName, ownerID)
		if err != nil {
			t.Errorf("unexpected error creating location: %v", err)
		}

		t.Log(got.ID)
	})

	t.Run("Create a scene", func(t *testing.T) {
		got, err := psqlRepo.CreateScene(context.Background(), projectName, locationName, ownerID, map[string]string{
			"A1": "ADD 15",
			"A2": "NEXT 66",
		})
		if err != nil {
			t.Errorf("failed to create a valid scene: %v", err)
		}
		t.Logf("got json: %v", got.Options)

		sceneID = got.ID
	})
	t.Run("Create another scene", func(t *testing.T) {
		got, err := psqlRepo.CreateScene(context.Background(), projectName, locationName, ownerID, map[string]string{
			"1": "ADD 15",
			"2": "NEXT 66",
		})
		if err != nil {
			t.Errorf("failed to create a valid scene: %v", err)
		}
		t.Logf("got json: %v", got.Options)
	})

	t.Run("Get all scenes in created location", func(t *testing.T) {
		got, err := psqlRepo.GetLocationScenes(context.Background(), projectName, locationName, ownerID)
		if err != nil {
			t.Errorf("failed to retrieve location scenes: %v", err)
		}

		t.Logf("got locs: %v", got)

	})

	t.Run("Get ownerID of a scene using utility function", func(t *testing.T) {
		got, err := psqlRepo.getSceneOwnerID(context.Background(), projectName, locationName, sceneID)
		if err != nil {
			t.Errorf("failed to acquire scene's ownerID: %v", err)
		}

		t.Logf("owner id: %v", got)
	})
	t.Run("Try to get ownerID passing incorrect projectName", func(t *testing.T) {
		_, err := psqlRepo.getSceneOwnerID(context.Background(), "incorrect", locationName, sceneID)
		if err == nil {
			t.Errorf("didn't raise error for non-existing project")
		}

	})

	t.Run("Delete a scene", func(t *testing.T) {
		err := psqlRepo.DeleteScene(context.Background(), projectName, locationName, ownerID, sceneID)
		if err != nil {
			t.Errorf("unexpected error trying to delete a scene")
		}
	})
	t.Run("try to delete a scene with invalid userID", func(t *testing.T) {
		err := psqlRepo.DeleteScene(context.Background(), projectName, locationName, "1337", sceneID)
		if err == nil {
			t.Errorf("expected to raise error")
		}
	})

	t.Cleanup(func() { removeRows(psqlDB, projectID) })
}

func removeRows(db *sql.DB, testProjID int32) {
	stmt := t.Projects.
		DELETE().
		WHERE(
			t.Projects.ID.EQ(j.Int32((testProjID))),
		)
	stmt.Exec(db)
}
