package dev_psql_adapter

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

// func TestDevAdapter(t *testing.T) {
// 	psqlDB, err := sql.Open("postgres",
// 		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 			host, port, user, password, dbname))
// 	if err != nil {
// 		log.Fatalf("failed to connect to db for testing: %q", err)
// 	}

// 	psqlRepo := NewDevPSQLRepository(psqlDB)

// 	t.Run("Valid create new project", func(t *testing.T) {
// 		ctx := context.Background()
// 		projectName := "test"
// 		ownerID := "47"
// 		_, err := psqlRepo.CreateProject(ctx, projectName, ownerID)
// 		if err != nil {
// 			t.Errorf("undexpected error: %v", err)

// 		}
// 	})
// }
