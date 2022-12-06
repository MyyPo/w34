package auth

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	authv1 "github.com/MyyPo/w34.Go/gen/go/auth/v1"
	"github.com/MyyPo/w34.Go/gen/psql/auth/public/model"
	t "github.com/MyyPo/w34.Go/gen/psql/auth/public/table"
	"github.com/MyyPo/w34.Go/internal/adapters/psql"
	"github.com/MyyPo/w34.Go/internal/pkg/validators"
	. "github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "host.docker.internal"
	port     = 1234
	user     = "spuser"
	password = "SPuser96"
	dbname   = "postgres"
)

func TestSignUp(t *testing.T) {
	psqlImpl := setupPsql(t)

	t.Run("Successful signup", func(t *testing.T) {

		req := &authv1.SignUpRequest{
			Username: "stubhello",
			Email:    "stubhello@stub.com",
			Password: "stubhello",
		}

		got, err := psqlImpl.SignUp(context.Background(), req)
		if err != nil {
			t.Errorf("unexpected error while trying to sign up: %q", err)
		}
		want := model.Accounts{
			Username: "stubhello",
		}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Try to signup with the taken username", func(t *testing.T) {
		req := &authv1.SignUpRequest{
			Username: "stubhello",
			Email:    "validemail@stub.com",
			Password: "stubhello",
		}

		_, err := psqlImpl.SignUp(context.Background(), req)
		if err == nil {
			t.Errorf("succesfully signed up with the taken username")
		}
	})
}

func removeRows(db *sql.DB) {
	stmt := t.Accounts.
		DELETE().
		WHERE(
			t.Accounts.Username.NOT_EQ(String("")),
		)
	stmt.Exec(db)
}

func setupPsql(t *testing.T) *AuthServer {
	psqlDB, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname))
	if err != nil {
		log.Fatalf("failed to connect to db for testing: %q", err)
	}
	psqlRepo := psql_adapters.NewPSQLRepository(psqlDB)
	authValidator, err := validators.NewAuthValidator(60, "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$")
	if err != nil {
		log.Fatalf("failed to initialize validator for testing: %q", err)
	}
	// remove all affected database rows after the tests
	t.Cleanup(func() { removeRows(psqlDB) })
	return NewAuthServer(psqlRepo, *authValidator)
}
