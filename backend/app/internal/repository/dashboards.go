package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Model struct {
	Id          int    `json:"id"`
	DashId      string `json:"dash_id"`
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type TypeModel struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

type DashboardsRepository struct {
	Db *sql.DB
}

func (d *DashboardsRepository) Get() ([]*Model, error) {
	query := `select * from dashboards`

	rows, err := d.Db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}

	models := []*Model{}
	for rows.Next() {
		var model *Model
		err = rows.Scan(&model.Id, &model.DashId, &model.Title, &model.Description, &model.CreatedAt, &model.UpdatedAt, &model.DeletedAt)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}

func (d *DashboardsRepository) GetById(id int) (*Model, error) {
	query := `select * from dashboards where id = $1`
	params := []any{id}

	row := d.Db.QueryRowContext(context.Background(), query, params...)

	if row.Err() != nil {
		if !errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, nil
		}
		return nil, row.Err()
	}

	var model *Model
	err := row.Scan(&model.Id, &model.DashId, &model.Title, &model.Description, &model.CreatedAt, &model.UpdatedAt, &model.DeletedAt)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (d *DashboardsRepository) GetAvailableTypes() ([]*TypeModel, error) {
	query := `select * from dashboard_item_types`

	rows, err := d.Db.QueryContext(context.Background(), query)

	if err != nil {
		return nil, err
	}

	var models []*TypeModel

	for rows.Next() {
		model := TypeModel{}
		err = rows.Scan(&model.Id, &model.Title, &model.Name)
		if err != nil {
			return nil, err
		}
		models = append(models, &model)
	}

	return models, nil
}

func (d *DashboardsRepository) GetByDashId(id string) (*Model, error) {
	query := `select * from dashboards where dash_id = $1`
	params := []any{id}

	row := d.Db.QueryRowContext(context.Background(), query, params...)

	if row.Err() != nil {
		if !errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, nil
		}
		return nil, row.Err()
	}

	var model *Model
	err := row.Scan(&model.Id, &model.DashId, &model.Title, &model.Description, &model.CreatedAt, &model.UpdatedAt, &model.DeletedAt)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (d *DashboardsRepository) Create(dto Model) (*Model, error) {
	query := `
		insert into dashboards (title, description)
		values ($1, $2)
	`
	params := []any{dto.Title, dto.Description}

	res, err := d.Db.ExecContext(context.Background(), query, params...)

	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	row, err := d.GetById(int(lastId))
	if err != nil {
		return nil, err
	}

	return row, nil
}
