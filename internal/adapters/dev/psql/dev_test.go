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

	psqlDB, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname))
	if err != nil {
		log.Fatalf("failed to connect to db for testing: %q", err)
	}

	psqlRepo := NewDevPSQLRepository(psqlDB)

	t.Run("Valid create new project", func(t *testing.T) {
		ownerID := "47"
		got, err := psqlRepo.CreateProject(context.Background(), projectName, ownerID)
		if err != nil {
			t.Errorf("undexpected error: %v", err)

		}

		projectID = got.ID
	})

	t.Run("Valid create a new location", func(t *testing.T) {
		got, err := psqlRepo.CreateLocation(context.Background(), projectID, locationName)
		if err != nil {
			t.Errorf("unexpected error creating location: %v", err)
		}

		t.Log(got.ID)
	})
	removeRows(psqlDB, 47, locationName)
}

func removeRows(db *sql.DB, ownerID int32, locationName string) {
	stmt := t.Projects.
		DELETE().
		WHERE(
			t.Projects.OwnerID.EQ(j.Int(int64(ownerID))),
		)
	stmt.Exec(db)

	stmt2 := t.Locations.
		DELETE().
		WHERE(
			t.Locations.Name.EQ(j.String(locationName)),
		)
	stmt2.Exec(db)

}
