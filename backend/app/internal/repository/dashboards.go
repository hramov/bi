package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
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

type Item struct {
	Id          string         `json:"id"`
	DashId      string         `json:"dash_id"`
	ItemType    int            `json:"type"`
	Position    sql.NullString `json:"position"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Options     string         `json:"rawOptions"`
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

func (d *DashboardsRepository) CreateItem(dto Item) (*int, error) {
	query := `
		insert into dashboard_items (dash_id, item_type, title, description, options)
		values ($1, $2, $3, $4, $5)
		returning id
	`

	fmt.Println(dto)

	newId, _ := uuid.NewUUID()
	params := []any{newId.String(), dto.ItemType, dto.Title, dto.Description, dto.Options}

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
		dash_id = $1, item_type = $2, title = $3, description = $4, options = $5
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
