package dashboard_repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/adapter/postgresrepo/dashboard/model"
	dashboard_entity "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/entity"
)

func (d *RepositoryImpl) GetByDashId(ctx context.Context, id string) (*dashboard_entity.Dashboard, error) {
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

	row := d.db.QueryRowContext(ctx, query, params...)

	if row.Err() != nil {
		if !errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, row.Err()
	}

	m := &dashboard_model.Model{}

	var rawItems []byte
	err := row.Scan(&m.Id, &m.DashId, &m.Title, &m.Description, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt, &rawItems)
	if err != nil {
		return nil, err
	}

	items := []*dashboard_entity.Item{}

	err = json.Unmarshal(rawItems, &items)
	if err != nil {
		return nil, err
	}

	dash := &dashboard_entity.Dashboard{
		Id:          m.Id,
		DashId:      m.DashId,
		Title:       m.Title,
		Description: m.Description.String,
		CreatedAt:   m.CreatedAt,
		Items:       items,
	}

	if m.UpdatedAt.Valid {
		dash.UpdatedAt = m.UpdatedAt.Time
	}

	if m.DeletedAt.Valid {
		dash.DeletedAt = m.DeletedAt.Time
	}

	return dash, nil
}
