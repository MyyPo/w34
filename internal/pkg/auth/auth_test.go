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
	"github.com/MyyPo/w34.Go/internal/pkg/auth/jwt"
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
	psqlImpl := setupPsqlRedis(t)

	var signInRefreshToken string
	var refreshedRefrToken string

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
		// t.Logf("Signup token: %s", res.GetTokens().GetRefreshToken())

	})
	time.Sleep(1 * time.Second)
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
	time.Sleep(1 * time.Second)
	t.Run("Try to signin with created account", func(t *testing.T) {
		req := &authv1.SignInRequest{
			UnOrEmail: "stubhello",
			Password:  "stubhelloe21eqw121",
		}
		res, err := psqlImpl.SignIn(context.Background(), req)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		// save the token for the following tests
		signInRefreshToken = res.GetTokens().GetRefreshToken()
		// t.Logf("signInToken: %s", signInRefreshToken)
		// t.Logf("signIn token: %s", signInRefreshToken)
	})
	time.Sleep(1 * time.Second)
	t.Run("Refresh the token", func(t *testing.T) {
		req := &authv1.RefreshTokensRequest{
			RefreshToken: signInRefreshToken,
		}
		// t.Logf("signInToken: %s", signInRefreshToken)
		res, err := psqlImpl.RefreshTokens(context.Background(), req)
		if err != nil {
			t.Errorf("refresh tokens error: %v", err)
		}
		refreshedRefrToken = res.GetTokens().GetRefreshToken()
		// t.Logf("token after refresh: %s", refreshedRefrToken)
	})
	time.Sleep(1 * time.Second)
	t.Run("Try to refresh token outside of db (one generated with sign in test)", func(t *testing.T) {
		req := &authv1.RefreshTokensRequest{
			RefreshToken: signInRefreshToken,
		}
		// t.Logf("signInToken: %s", signInRefreshToken)
		// res, err := psqlImpl.RefreshTokens(context.Background(), req)
		_, err := psqlImpl.RefreshTokens(context.Background(), req)
		if err == nil {
			t.Errorf("expected error while trying to refresh with an old token")
		}
		// t.Logf(res.GetTokens().GetRefreshToken())
	})
	time.Sleep(1 * time.Second)
	t.Run("Refresh with a token acquired from refresh method", func(t *testing.T) {
		req := &authv1.RefreshTokensRequest{
			RefreshToken: refreshedRefrToken,
		}
		_, err := psqlImpl.RefreshTokens(context.Background(), req)
		if err != nil {
			t.Errorf("refresh tokens error: %v", err)
		}
		// t.Logf(res.GetTokens().GetRefreshToken())

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

func setupPsqlRedis(t *testing.T) *AuthServer {
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
	hasher := hasher.NewHasher()
	psqlRepo := auth_psql_adapter.NewPSQLRepository(psqlDB)
	// redisClient := auth_redis.NewRedisClient("localhost:6379", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81")
	redisClient := auth_redis.NewRedisClient("host.docker.internal:6379", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81")
	authValidator, err := validators.NewAuthValidator(60, "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$")
	if err != nil {
		log.Fatalf("failed to initialize validator for testing: %q", err)
	}

	jwtManager, err := jwt.NewJWTManager("../../../configs/rsa", "../../../configs/rsa.pub",
		"../../../configs/refresh_rsa", "../../../configs/refresh_rsa.pub",
		config.AccessTokenDuration, config.RefreshTokenDuration)
	if err != nil {
		log.Fatalf("failed to initialize jwtManager for testing: %q", err)
	}
	// remove all affected database rows after the tests
	t.Cleanup(func() { removeRows(psqlDB) })
	return NewAuthServer(psqlRepo, *redisClient, *authValidator, *jwtManager, *hasher)
}
