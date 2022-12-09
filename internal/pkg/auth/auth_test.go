package auth

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/MyyPo/w34.Go/configs"
	authv1 "github.com/MyyPo/w34.Go/gen/go/auth/v1"
	t "github.com/MyyPo/w34.Go/gen/psql/auth/public/table"
	"github.com/MyyPo/w34.Go/internal/adapters/auth/psql"
	"github.com/MyyPo/w34.Go/internal/pkg/auth/hasher"
	auth_redis "github.com/MyyPo/w34.Go/internal/pkg/auth/redis"
	"github.com/MyyPo/w34.Go/internal/pkg/auth/validators"
	. "github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "host.docker.internal"
	port     = 1234
	user     = "spuser"
	password = "SPuser96"
	// dbname   = "postgres"
	dbname = "auth"
)

func TestSignUpSignIn(t *testing.T) {
	psqlImpl := setupPsql(t)
	_, err := configs.NewConfig("../../../configs")
	if err != nil {
		t.Errorf("failed to load config: %q", err)
	}

	t.Run("Successful signup", func(t *testing.T) {

		req := &authv1.SignUpRequest{
			Username: "stubhello",
			Email:    "stubhello@stub.com",
			Password: "stubhelloe21eqw121",
		}

		_, err := psqlImpl.SignUp(context.Background(), req)
		if err != nil {
			t.Errorf("unexpected error while trying to sign up: %q", err)
		}
	})
	t.Run("Try to signup with the taken username", func(t *testing.T) {
		req := &authv1.SignUpRequest{
			Username: "stubhello",
			Email:    "validemail@stub.com",
			Password: "stubhelloe21eqw121",
		}

		_, err := psqlImpl.SignUp(context.Background(), req)
		if err == nil {
			t.Errorf("succesfully signed up with the taken username")
		}
	})
	t.Run("Try to signin with created account", func(t *testing.T) {
		req := &authv1.SignInRequest{
			UnOrEmail: "stubhello",
			Password:  "stubhelloe21eqw121",
		}
		_, err := psqlImpl.SignIn(context.Background(), req)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		// t.Logf("signing: %v", got.GetTokens().AccessToken)
	})
}

func removeRows(db *sql.DB) {
	stmt := t.Accounts.
		DELETE().
		WHERE(
			t.Accounts.Username.EQ(String("stubhello")),
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
	psqlRepo := auth_psql_adapter.NewPSQLRepository(psqlDB)
	hasher := hasher.NewHasher()
	redisClient := auth_redis.NewRedisClient("localhost:6379", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", *hasher)
	authValidator, err := validators.NewAuthValidator(60, "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$")
	if err != nil {
		log.Fatalf("failed to initialize validator for testing: %q", err)
	}

	jwtManager := NewJWTManager("../../../configs/rsa", "../../../configs/rsa.pub",
		"../../../configs/refresh_rsa", "../../../configs/refresh_rsa.pub",
		time.Minute*10, time.Hour*48)

	// remove all affected database rows after the tests
	t.Cleanup(func() { removeRows(psqlDB) })
	return NewAuthServer(psqlRepo, *redisClient, *authValidator, *jwtManager)
}
