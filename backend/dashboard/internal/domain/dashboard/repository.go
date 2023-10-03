package dashboard

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	dashboards_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/in"
)

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (d *RepositoryImpl) Get() ([]*Dashboard, error) {
	query := `select * from dashboards`

	rows, err := d.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}

	dashboards := []*Dashboard{}

	for rows.Next() {
		model := new(Model)
		err = rows.Scan(&model.Id, &model.DashId, &model.Title, &model.Description, &model.CreatedAt, &model.UpdatedAt, &model.DeletedAt)
		if err != nil {
			return nil, err
		}

		dash := &Dashboard{
			Id:          model.Id,
			DashId:      model.DashId,
			Title:       model.Title,
			Description: model.Description.String,
			CreatedAt:   model.CreatedAt,
			Items:       nil,
		}

		if model.UpdatedAt.Valid {
			dash.UpdatedAt = model.UpdatedAt.Time
		}

		if model.DeletedAt.Valid {
			dash.DeletedAt = model.DeletedAt.Time
		}

		dashboards = append(dashboards, dash)
	}

	return dashboards, nil
}

func (d *RepositoryImpl) GetItemById(id int) (*Item, error) {
	query := `select * from dashboard_items where id = $1`
	params := []any{id}

	row := d.db.QueryRowContext(context.Background(), query, params...)

	if row.Err() != nil {
		if !errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, nil
		}
		return nil, row.Err()
	}

	model := new(Item)
	err := row.Scan(&model.DashId, &model.ItemType, &model.Position, &model.Title, &model.Description, &model.Options, &model.CreatedAt, &model.UpdatedAt, &model.DeletedAt, &model.Id)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (d *RepositoryImpl) GetAvailableTypes() ([]*ItemType, error) {
	query := `select * from dashboard_item_types`

	rows, err := d.db.QueryContext(context.Background(), query)

	if err != nil {
		return nil, err
	}

	types := []*ItemType{}

	for rows.Next() {
		model := ItemTypeModel{}
		err = rows.Scan(&model.Id, &model.Title, &model.Name)
		if err != nil {
			return nil, err
		}

		t := &ItemType{
			Id:    model.Id,
			Title: model.Title,
			Name:  model.Name,
		}

		types = append(types, t)
	}

	return types, nil
}

func (d *RepositoryImpl) GetByDashId(id string) (*Dashboard, error) {
	query := `
		select d.id, d.dash_id, d.title, d.description,
		   date_trunc('second', d.created_at) as created_at,
		   date_trunc('second', d.updated_at) as updated_at,
		   date_trunc('second', d.deleted_at) as deleted_at,
		   COALESCE(jsonb_agg(
			   json_build_object(
						'id', di.id, 'dash_id', di.dash_id, 'item_type', di.item_type, 'position', di.position, 'title', di.title,
						'description', di.description, 'data_queries', di.data_queries, 'options', di.raw_options, 'created_at', to_char(di.created_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"'),
						'updated_at', to_char(di.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"'), 'deleted_at', to_char(di.deleted_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"')
				   )
			   ) FILTER (WHERE di.dash_id IS NOT NULL), '[]') AS items
		from dashboards d
		left join dashboard_items di on di.dash_id = d.dash_id
		where d.dash_id = $1
		group by d.id, d.dash_id, d.title, d.description, d.created_at, d.updated_at, d.deleted_at
	`
	params := []any{id}

	row := d.db.QueryRowContext(context.Background(), query, params...)

	if row.Err() != nil {
		if !errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, nil
		}
		return nil, row.Err()
	}

	model := &Model{}

	var rawItems []byte
	err := row.Scan(&model.Id, &model.DashId, &model.Title, &model.Description, &model.CreatedAt, &model.UpdatedAt, &model.DeletedAt, &rawItems)
	if err != nil {
		return nil, err
	}

	items := []*Item{}

	err = json.Unmarshal(rawItems, &items)
	if err != nil {
		return nil, err
	}

	dash := &Dashboard{
		Id:          model.Id,
		DashId:      model.DashId,
		Title:       model.Title,
		Description: model.Description.String,
		CreatedAt:   model.CreatedAt,
		Items:       items,
	}

	if model.UpdatedAt.Valid {
		dash.UpdatedAt = model.UpdatedAt.Time
	}

	if model.DeletedAt.Valid {
		dash.DeletedAt = model.DeletedAt.Time
	}

	return dash, nil
}

func (d *RepositoryImpl) Create(dto dashboards_dto_in.Dashboard) (*int, error) {
	query := `
		insert into dashboards (title, description)
		values ($1, $2)
		returning id;
	`
	params := []any{dto.Title, dto.Description}

	row := d.db.QueryRowContext(context.Background(), query, params...)

	var id int
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}

func (d *RepositoryImpl) Update(dto dashboards_dto_in.Dashboard, id int) (*int, error) {
	query := `
		update dashboards set
		title = $1, description = $2, updated_at = now()
		where id = $3
		returning id
	`
	params := []any{dto.Title, dto.Description, id}

	row := d.db.QueryRowContext(context.Background(), query, params...)

	var updatedId int
	if err := row.Scan(&updatedId); err != nil {
		return nil, err
	}

	return &updatedId, nil
}

func (d *RepositoryImpl) CreateItem(dto dashboards_dto_in.Item) (*int, error) {
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

	row := d.db.QueryRowContext(context.Background(), query, params...)

	var id int
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}

func (d *RepositoryImpl) UpdateItem(dto dashboards_dto_in.Item, id int) (*int, error) {
	query := `
		update dashboard_items set
		dash_id = $1, item_type = $2, title = $3, description = $4, raw_options = $5, updated_at = now()
		where id = $6
		returning id
	`

	params := []any{dto.DashId, dto.ItemType, dto.Title, dto.Description, dto.Options, id}

	row := d.db.QueryRowContext(context.Background(), query, params...)

	var updatedId int

	if err := row.Scan(&updatedId); err != nil {
		return nil, err
	}

	return &updatedId, nil
}
