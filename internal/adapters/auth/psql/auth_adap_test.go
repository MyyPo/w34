package auth_psql_adapter

import (
	"context"
	"database/sql"
	"fmt"
	t "github.com/MyyPo/w34.Go/gen/psql/main/public/table"
	j "github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
	"log"
	"testing"
)

const (
	host     = "host.docker.internal"
	port     = 5432
	user     = "spuser"
	password = "SPuser96"
	dbname   = "auth"
)

func TestAuthAdapter(t *testing.T) {
	username := "authadapuser"
	email := "authadapemail@gmail.com"
	var testUserID int32

	psqlDB, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname))
	if err != nil {
		log.Fatalf("failed to connect to db for testing: %q", err)
	}
	psqlRepo := NewAuthPSQLRepository(psqlDB)

	t.Run("Create a new user", func(t *testing.T) {
		res, err := psqlRepo.CreateUser(context.Background(), username, email, "password")
		if err != nil {
			t.Errorf("failed to create a new user: %v", err)
		}

		testUserID = res.UserID
	})

	t.Run("Lookup existing user by username", func(t *testing.T) {
		_, err := psqlRepo.LookupExistingUser(context.Background(), username)
		if err != nil {
			t.Errorf("undexpected error: %v", err)
		}
	})
	t.Run("Lookup existing user by email", func(t *testing.T) {
		_, err := psqlRepo.LookupExistingUser(context.Background(), email)
		if err != nil {
			t.Errorf("undexpected error: %v", err)
		}
	})

	t.Cleanup(func() { removeRows(psqlDB, testUserID) })
}

func removeRows(db *sql.DB, testUserID int32) {
	stmt := t.Accounts.
		DELETE().
		WHERE(
			t.Accounts.UserID.EQ(j.Int32((testUserID))),
		)
	stmt.Exec(db)
}
