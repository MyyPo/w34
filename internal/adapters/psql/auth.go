package psql_adapters

import (
	"context"
	"database/sql"

	authv1 "github.com/MyyPo/w34.Go/gen/go/auth/v1"
	"github.com/MyyPo/w34.Go/gen/psql/auth/public/model"
	. "github.com/MyyPo/w34.Go/gen/psql/auth/public/table"
	// . "github.com/go-jet/jet/v2/postgres"
)

type PSQLRepository struct {
	db *sql.DB
}

func NewPSQLRepository(db *sql.DB) *PSQLRepository {
	return &PSQLRepository{db: db}
}

func (r PSQLRepository) CreateUser(ctx context.Context, req *authv1.SignUpRequest) (model.Accounts, error) {
	credentials := req.GetCredentials()

	stmt := Accounts.
		INSERT(
			Accounts.Username,
			Accounts.Password,
		).VALUES(
		credentials.GetUsername(),
		credentials.GetPassword(),
	).RETURNING(
		Accounts.UserID,
		Accounts.Username,
	)

	var result model.Accounts
	err := stmt.Query(r.db, &result)
	if err != nil {
		return model.Accounts{}, err
	}

	return result, nil
}
