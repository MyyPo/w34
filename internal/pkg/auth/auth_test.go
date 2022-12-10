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
	psqlImpl := setupPsqlRedis(t)
	_, err := configs.NewConfig("../../../configs")
	if err != nil {
		t.Errorf("failed to load config: %q", err)
	}

	var finalTestRefreshToken string

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
		res, err := psqlImpl.SignIn(context.Background(), req)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		// save tokens for the next test
		finalTestRefreshToken = res.GetTokens().GetRefreshToken()
	})
	t.Run("Refresh the token", func(t *testing.T) {
		req := &authv1.RefreshTokensRequest{
			RefreshToken: finalTestRefreshToken,
		}
		res, err := psqlImpl.RefreshTokens(context.Background(), req)
		if err != nil {
			t.Errorf("refresh tokens error: %v", err)
		}
		endAccessToken := res.GetTokens().GetAccessToken()
		t.Logf("end token: %s", endAccessToken)
	})
	t.Run("Try to refresh token outside of db", func(t *testing.T) {
		const oldRefreshToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0a25fdHlwZSI6InJlZnJlc2giLCJzdWIiOiI3NCIsImV4cCI6MTY3MDg0MDg4MywibmJmIjoxNjcwNjY4MDgzLCJpYXQiOjE2NzA2NjgwODN9.s7BxdTQ_BiF2Yd8IJhC-WFc92LT06Fp8d3UiYT31_Cn8wgn0QIIYguNlJq_Jtt9teYG5cddTq5bKys_Oxx56HEdcusyFBlxW8qsoCGQKDjvE1a9Se7E1uIDr53uTQcv6Jk-E_KONXOXyslxvm3jlMZz5qpNaiR0miZV9IyqxqLysbHdzxjd7YH-El53TUFKKmJJjQ0jpLIhI5jEwMB3W-O8vMOr21QqnQfToGpGJCjmVFKT9g-truoiRpyZUmnpFdj6bftcn9KU5UdQ2g1if8G67OUor7IKN0Ra63A5YeW-0D8mUwk_EX7qnZOG5BSX4bRvaxJH42ccqzCkhqXYr-speQ68yspdb8-nAZ1TIaOB-8kxc26gWIs8SYExL-GToVlvXd1hQLWKQ0ZlWeI4mFSSxa34V5R5x2voLvpDO5IEz3eeF9yuyr-fgt1zIpt75z0h5Zb0eH_XamnOdcYuOgNNBiTNseUaaOCW6uhElaBjFGfdsEuyobeWX09u0cyvh4DQ4bQKc42AvshCt5KmdCPUjtvZkR6ixclT7b9dY6Hjux2gOOUn1RM_8iYSSGNpB57vTbCTKPwcvoS-JlmfBv-buayMCAjWyTpImkLqNHK0XNyyNusCFs6EncQXp-Z5f86aCHBi7iSqqVUvR3CBKJQ_C580t1WMFjKSF61bzZz4"
		req := &authv1.RefreshTokensRequest{
			RefreshToken: oldRefreshToken,
		}
		_, err := psqlImpl.RefreshTokens(context.Background(), req)
		if err == nil {
			t.Errorf("expected error while trying to refresh with an old token")
		}

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
