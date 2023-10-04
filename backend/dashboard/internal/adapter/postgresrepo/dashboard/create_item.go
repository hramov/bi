package dashboard_repo

import (
	"context"
	"encoding/json"
	dashboards_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/in"
)

func (d *RepositoryImpl) CreateItem(ctx context.Context, dto dashboards_dto_in.Item) (*int, error) {
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

	row := d.db.QueryRowContext(ctx, query, params...)

	var id int
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}
