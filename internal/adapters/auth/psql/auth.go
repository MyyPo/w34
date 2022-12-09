package auth_psql_adapter

import (
	"context"
	"database/sql"

	"github.com/MyyPo/w34.Go/gen/psql/auth/public/model"
	t "github.com/MyyPo/w34.Go/gen/psql/auth/public/table"
	j "github.com/go-jet/jet/v2/postgres"
)

type PSQLRepository struct {
	db *sql.DB
}

func NewPSQLRepository(db *sql.DB) *PSQLRepository {
	return &PSQLRepository{db: db}
}

func (r PSQLRepository) CreateUser(
	ctx context.Context,
	newUsername string,
	newEmail string,
	newHashedPassword string,
) (model.Accounts, error) {

	stmt := t.Accounts.
		INSERT(
			t.Accounts.Username,
			t.Accounts.Email,
			t.Accounts.Password,
		).VALUES(
		newUsername,
		newEmail,
		newHashedPassword,
	).RETURNING(
		t.Accounts.UserID,
		t.Accounts.Username,
	)

	var result model.Accounts
	err := stmt.Query(r.db, &result)
	if err != nil {
		return model.Accounts{}, err
	}

	return result, nil
}

func (r PSQLRepository) LookupExistingUser(
	ctx context.Context,
	usernameOrEmail string,
) (model.Accounts, error) {

	stmt := t.Accounts.
		SELECT(
			t.Accounts.UserID,
			t.Accounts.Password,
		).WHERE(
		t.Accounts.Username.EQ(j.String(usernameOrEmail)).
			OR(t.Accounts.Email.EQ(j.String(usernameOrEmail))),
	)

	var result model.Accounts
	err := stmt.Query(r.db, &result)
	if err != nil {
		return model.Accounts{}, err
	}

	return result, nil
}
