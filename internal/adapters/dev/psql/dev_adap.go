package dev_psql_adapter

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/MyyPo/w34.Go/gen/psql/main/public/model"
	t "github.com/MyyPo/w34.Go/gen/psql/main/public/table"
	j "github.com/go-jet/jet/v2/postgres"
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
		t.Projects.ID,
		t.Projects.Name,
	)

	var result model.Projects
	err := stmt.Query(r.db, &result)
	if err != nil {
		return model.Projects{}, err
	}

	return result, nil
}

func (r DevPSQLRepository) DeleteProject(
	ctx context.Context,
	projectName string,
	reqUserID string,
) error {
	intReqUserID, err := strconv.ParseInt(reqUserID, 10, 32)
	if err != nil {
		return fmt.Errorf("internal error")
	}

	stmt := t.Projects.
		DELETE().
		WHERE(
			t.Projects.OwnerID.EQ(j.Int(intReqUserID)).
				AND(
					t.Projects.Name.EQ(j.String(projectName))),
		)

	res, err := stmt.Exec(r.db)

	if err != nil {
		return err
	}

	rowsDeleted, _ := res.RowsAffected()
	if rowsDeleted == 0 {
		return fmt.Errorf("there is no project with such name")
	}

	return nil
}

func (r DevPSQLRepository) CreateLocation(
	ctx context.Context,
	projectName string,
	locationName string,
	reqUserID string,
) (model.Locations, error) {
	intReqUserID, err := strconv.ParseInt(reqUserID, 10, 32)
	if err != nil {
		return model.Locations{}, fmt.Errorf("internal error")
	}

	lookupProjectID := t.Projects.
		SELECT(
			t.Projects.ID,
		).WHERE(
		t.Projects.OwnerID.EQ(j.Int(intReqUserID)).
			AND(
				t.Projects.Name.EQ(j.String(projectName)),
			),
	)

	var lookupResult model.Projects
	err = lookupProjectID.Query(r.db, &lookupResult)
	if err != nil {
		return model.Locations{}, err
	}

	stmt := t.Locations.
		INSERT(
			t.Locations.ProjectID,
			t.Locations.Name,
		).
		VALUES(
			lookupResult.ID,
			locationName,
		).RETURNING(
		t.Locations.ID,
	)
	var result model.Locations
	err = stmt.Query(r.db, &result)
	if err != nil {
		return model.Locations{}, err
	}

	return result, nil
}

func (r DevPSQLRepository) CreateScene(
	ctx context.Context,
	projectName string,
	locationName string,
	reqUserID string,
	sceneOptions map[string]string,
) (model.Scenes, error) {
	lookupLocation, err := r.getLocationID(projectName, locationName, reqUserID)
	if err != nil {
		return model.Scenes{}, err
	}

	jsonSceneOptions, err := json.Marshal(sceneOptions)
	if err != nil {
		return model.Scenes{}, err
	}

	stmt := t.Scenes.INSERT(
		t.Scenes.LocationID,
		t.Scenes.Options,
	).VALUES(
		lookupLocation.ID,
		jsonSceneOptions,
	).RETURNING(
		t.Scenes.ID,
		t.Scenes.Options,
	)

	var result model.Scenes
	err = stmt.Query(r.db, &result)
	if err != nil {
		return model.Scenes{}, err
	}

	return result, nil
}

func (r DevPSQLRepository) DeleteScene(
	ctx context.Context,
	projectName string,
	locationName string,
	reqUserID string,
) (model.Scenes, error) {
	return model.Scenes{}, nil
}

func (r DevPSQLRepository) GetLocationScenes(
	ctx context.Context,
	projectName string,
	locationName string,
	reqUserID string,
) ([]model.Scenes, error) {
	lookupLocation, err := r.getLocationID(projectName, locationName, reqUserID)
	if err != nil {
		return []model.Scenes{}, err
	}

	stmt := t.Scenes.SELECT(
		t.Scenes.ID,
		t.Scenes.Options,
	).WHERE(
		t.Scenes.LocationID.EQ(j.Int(int64(lookupLocation.ID))),
	)
	var result []model.Scenes
	err = stmt.Query(r.db, &result)
	if err != nil {
		return []model.Scenes{}, err
	}

	return result, nil
}

// Util method giving access to location id by name
func (r DevPSQLRepository) getLocationID(
	projectName string,
	locationName string,
	reqUserID string,
) (model.Locations, error) {
	intReqUserID, err := strconv.ParseInt(reqUserID, 10, 32)
	if err != nil {
		return model.Locations{}, fmt.Errorf("internal error")
	}

	stmt := j.SELECT(
		t.Locations.ID,
	).FROM(
		t.Locations.
			INNER_JOIN(t.Projects, t.Projects.Name.EQ(j.String(projectName)).
				AND(t.Projects.OwnerID.EQ(j.Int(intReqUserID))).
				AND(t.Locations.Name.EQ(j.String(locationName))),
			),
	)
	var lookupResult model.Locations
	err = stmt.Query(r.db, &lookupResult)
	if err != nil {
		return model.Locations{}, err
	}

	return lookupResult, err
}
