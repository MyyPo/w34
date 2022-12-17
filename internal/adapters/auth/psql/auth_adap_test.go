package auth_psql_adapter

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	_ "github.com/lib/pq"
// 	"log"
// 	"testing"
// )

const (
	host     = "host.docker.internal"
	port     = 5432
	user     = "spuser"
	password = "SPuser96"
	dbname   = "auth"
)

// func TestAuthAdapter(t *testing.T) {
// 	psqlDB, err := sql.Open("postgres",
// 		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 			host, port, user, password, dbname))
// 	if err != nil {
// 		log.Fatalf("failed to connect to db for testing: %q", err)
// 	}
// 	psqlRepo := NewPSQLRepository(psqlDB)

// 	t.Run("Lookup existing user by username", func(t *testing.T) {
// 		ctx := context.Background()
// 		username := "test"
// 		_, err := psqlRepo.LookupExistingUser(ctx, username)
// 		if err != nil {
// 			t.Errorf("undexpected error: %v", err)
// 		}
// 	})
// 	t.Run("Lookup existing user by email", func(t *testing.T) {
// 		ctx := context.Background()
// 		email := "test@test.com"
// 		_, err := psqlRepo.LookupExistingUser(ctx, email)
// 		if err != nil {
// 			t.Errorf("undexpected error: %v", err)
// 		}
// 	})
// }
