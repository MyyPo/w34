package auth_psql_adapter

import (
	"context"
	"database/sql"

	"github.com/MyyPo/w34.Go/gen/psql/main/public/model"
	t "github.com/MyyPo/w34.Go/gen/psql/main/public/table"
	j "github.com/go-jet/jet/v2/postgres"
)

type AuthPSQLRepository struct {
	db *sql.DB
}

func NewAuthPSQLRepository(db *sql.DB) *AuthPSQLRepository {
	return &AuthPSQLRepository{db: db}
}

func (r AuthPSQLRepository) CreateUser(
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

func (r AuthPSQLRepository) LookupExistingUser(
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
