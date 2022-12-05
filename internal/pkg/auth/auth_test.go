package auth

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	authv1 "github.com/MyyPo/w34.Go/gen/go/auth/v1"
	"github.com/MyyPo/w34.Go/gen/psql/auth/public/model"
	"github.com/MyyPo/w34.Go/internal/adapters/psql"
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
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname))
	if err != nil {
		t.Errorf("failed to connect to db for testing: %q", err)
	}

	t.Run("Saying heeey", func(t *testing.T) {
		rep := psql_adapters.NewPSQLRepository(db)
		impl := NewAuthServer(rep)

		req := &authv1.SignUpRequest{
			Credentials: &authv1.AuthRequest{
				Username: "stubhello",
				Password: "stubhello",
			},
		}

		got, err := impl.SignUp(context.Background(), req)
		if err != nil {
			t.Errorf("unexpected error: %q", err)
		}
		want := model.Accounts{
			Username: "stubhello",
		}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
