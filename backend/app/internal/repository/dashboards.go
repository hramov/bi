package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

type Model struct {
	Id          int            `json:"id"`
	DashId      string         `json:"dash_id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	DeletedAt   sql.NullTime   `json:"deleted_at"`

	Items string `json:"items"`
}

type TypeModel struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

type Item struct {
	Id          int            `json:"id"`
	DashId      string         `json:"dash_id"`
	ItemType    int            `json:"type"`
	Position    sql.NullString `json:"position"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Options     any            `json:"raw_options"`
	RawOptions  any            `json:"options"`
	DataQueries any            `json:"data_queries"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	DeletedAt   sql.NullTime   `json:"deleted_at"`
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
		model := &Model{}
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

func (d *DashboardsRepository) GetItemById(id int) (*Item, error) {
	query := `select * from dashboard_items where id = $1`
	params := []any{id}

	row := d.Db.QueryRowContext(context.Background(), query, params...)

	if row.Err() != nil {
		if !errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, nil
		}
		return nil, row.Err()
	}

	model := &Item{}
	err := row.Scan(&model.DashId, &model.ItemType, &model.Position, &model.Title, &model.Description, &model.Options, &model.CreatedAt, &model.UpdatedAt, &model.DeletedAt, &model.Id)
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
	query := `
		select d.id, d.dash_id, d.title, d.description, 
		       date_trunc('second', d.created_at) created_at, 
		       date_trunc('second', d.updated_at) updated_at, 
		       date_trunc('second', d.deleted_at) deleted_at,
			   COALESCE(json_agg(di) FILTER (WHERE di.dash_id IS NOT NULL), '[]') AS items
		from dashboards d
		left join dashboard_items di on di.dash_id = d.dash_id
		where d.dash_id = $1
		group by d.id, d.dash_id, d.title, d.description, d.created_at, d.updated_at, d.deleted_at
	`
	params := []any{id}

	row := d.Db.QueryRowContext(context.Background(), query, params...)

	if row.Err() != nil {
		if !errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, nil
		}
		return nil, row.Err()
	}

	model := &Model{}
	err := row.Scan(&model.Id, &model.DashId, &model.Title, &model.Description, &model.CreatedAt, &model.UpdatedAt, &model.DeletedAt, &model.Items)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (d *DashboardsRepository) Create(dto Model) (*int, error) {
	query := `
		insert into dashboards (title, description)
		values ($1, $2)
		returning id;
	`
	params := []any{dto.Title, dto.Description}

	row := d.Db.QueryRowContext(context.Background(), query, params...)

	var id int
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}

func (d *DashboardsRepository) Update(dto Model, id int) (*int, error) {
	query := `
		update dashboards set
		title = $1, description = $2, updated_at = now()
		where id = $3
		returning id
	`
	params := []any{dto.Title, dto.Description, id}

	row := d.Db.QueryRowContext(context.Background(), query, params...)

	var updatedId int
	if err := row.Scan(&updatedId); err != nil {
		return nil, err
	}

	return &updatedId, nil
}

func (d *DashboardsRepository) CreateItem(dto Item) (*int, error) {
	query := `
		insert into dashboard_items (dash_id, item_type, title, description, raw_options, data_queries)
		values ($1, $2, $3, $4, $5, $6)
		returning id
	`

	options, err := json.Marshal(dto.Options)
	if err != nil {
		return nil, err
	}

	dq, err := json.Marshal(dto.DataQueries)
	if err != nil {
		return nil, err
	}

	params := []any{dto.DashId, dto.ItemType, dto.Title, dto.Description, options, dq}

	row := d.Db.QueryRowContext(context.Background(), query, params...)

	var id int
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}

func (d *DashboardsRepository) UpdateItem(dto Item, id int) (*int, error) {
	query := `
		update dashboard_items set
		dash_id = $1, item_type = $2, title = $3, description = $4, raw_options = $5, updated_at = now()
		where id = $6
		returning id
	`

	params := []any{dto.DashId, dto.ItemType, dto.Title, dto.Description, dto.Options, id}

	row := d.Db.QueryRowContext(context.Background(), query, params...)

	var updatedId int

	if err := row.Scan(&updatedId); err != nil {
		return nil, err
	}

	return &updatedId, nil
}
