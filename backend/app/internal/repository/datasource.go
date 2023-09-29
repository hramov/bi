package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/hramov/gvc-bi/backend/internal/api/http/handlers/datasource"
)

type DatasourceRepository struct {
	Db *sql.DB
}

func (d *DatasourceRepository) GetDrivers() ([]*datasource.Driver, error) {
	query := `select * from drivers`

	rows, err := d.Db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}

	models := []*datasource.Driver{}

	for rows.Next() {
		model := &datasource.Driver{}
		err = rows.Scan(&model.Id, &model.Title, &model.Code, &model.DateCreated, &model.DbNeed)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}

func (d *DatasourceRepository) GetDriverById(id int) (*datasource.Driver, error) {
	query := `select * from drivers where id = $1`
	params := []any{id}

	row := d.Db.QueryRowContext(context.Background(), query, params...)

	if row.Err() != nil {
		if !errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, nil
		}
		return nil, row.Err()
	}

	model := &datasource.Driver{}
	err := row.Scan(&model.Id, &model.Title, &model.Code, &model.DateCreated, &model.DbNeed)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (d *DatasourceRepository) Get() ([]*datasource.Datasource, error) {
	query := `
		select ds.id, ds.driver_id, ds.title, ds.dsn, ds.checked, ds.date_created, d.code  
		from data_sources ds
		join drivers d on d.id = ds.driver_id
	`

	rows, err := d.Db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}

	models := []*datasource.Datasource{}

	for rows.Next() {
		model := &datasource.Datasource{}
		err = rows.Scan(&model.Id, &model.DriverId, &model.Title, &model.Dsn, &model.Checked, &model.DateCreated, &model.Driver)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}

func (d *DatasourceRepository) GetById(id int) (*datasource.Datasource, error) {
	query := `
		select ds.id, ds.driver_id, ds.title, ds.dsn, ds.checked, ds.date_created, d.code  
		from data_sources ds
		join drivers d on d.id = ds.driver_id
		where ds.id = $1
	`
	params := []any{id}

	row := d.Db.QueryRowContext(context.Background(), query, params...)

	if row.Err() != nil {
		if !errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, nil
		}
		return nil, row.Err()
	}

	model := &datasource.Datasource{}
	err := row.Scan(&model.Id, &model.DriverId, &model.Title, &model.Dsn, &model.Checked, &model.DateCreated, &model.Driver)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (d *DatasourceRepository) GetByCode(code string) (*datasource.Datasource, error) {
	query := `
		select ds.id, ds.driver_id, ds.title, ds.dsn, ds.checked, ds.date_created, d.code  
		from data_sources ds
		join drivers d on d.id = ds.driver_id
		where d.code = $1
	`
	params := []any{code}

	row := d.Db.QueryRowContext(context.Background(), query, params...)

	if row.Err() != nil {
		if !errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, nil
		}
		return nil, row.Err()
	}

	model := &datasource.Datasource{}
	err := row.Scan(&model.Id, &model.DriverId, &model.Title, &model.Dsn, &model.Checked, &model.DateCreated, &model.Driver)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (d *DatasourceRepository) Create(source datasource.Datasource) (*int, error) {
	query := `
		INSERT INTO data_sources (driver_id, title, dsn, checked)
		SELECT id, $2, $3, $4 from drivers WHERE code = $1
		RETURNING id
	`

	params := []any{source.Driver, source.Title, source.Dsn, source.Checked}

	row := d.Db.QueryRowContext(context.Background(), query, params...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var id int
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (d *DatasourceRepository) Update(source datasource.Datasource, id int) (*int, error) {
	query := `
		UPDATE data_sources
		SET driver_id = (SELECT id FROM drivers WHERE code = $1), title = $2, dsn = $3, checked = $4
		WHERE id = $5
		RETURNING id
	`

	params := []any{source.Driver, source.Title, source.Dsn, source.Checked, id}

	row := d.Db.QueryRowContext(context.Background(), query, params...)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var updatedId int
	err := row.Scan(&updatedId)
	if err != nil {
		return nil, err
	}

	return &updatedId, nil
}

func (d *DatasourceRepository) Delete(id int) (*int, error) {
	query := `DELETE FROM data_sources WHERE id = $1 RETURNING id`

	params := []any{id}

	row := d.Db.QueryRowContext(context.Background(), query, params...)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var deletedId int
	err := row.Scan(&deletedId)
	if err != nil {
		return nil, err
	}

	return &deletedId, nil
}
