package dev_psql_adapter

import (
	"context"
	"database/sql"

	"github.com/MyyPo/w34.Go/gen/psql/main/public/model"
	t "github.com/MyyPo/w34.Go/gen/psql/main/public/table"
)

type DevPSQLRepository struct {
	db *sql.DB
}

func NewDevPSQLRepository(db *sql.DB) *DevPSQLRepository {
	return &DevPSQLRepository{db: db}
}

func (r DevPSQLRepository) CreateProject(
	ctx context.Context,
	projectName string,
	ownerID string,
) (model.Projects, error) {

	stmt := t.Projects.
		INSERT(
			t.Projects.Name,
			t.Projects.OwnerID,
		).VALUES(
		projectName,
		ownerID,
	).RETURNING(
		t.Projects.Name,
	)

	var result model.Projects
	err := stmt.Query(r.db, &result)
	if err != nil {
		return model.Projects{}, err
	}

	return result, nil
}
