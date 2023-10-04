package dashboard_repo

import (
	"context"
	dashboards_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/in"
)

func (d *RepositoryImpl) UpdateItem(ctx context.Context, dto dashboards_dto_in.Item, id int) (*int, error) {
	query := `
		update dashboard_items set
		dash_id = $1, item_type = $2, title = $3, description = $4, raw_options = $5, updated_at = now()
		where id = $6
		returning id
	`

	params := []any{dto.DashId, dto.ItemType, dto.Title, dto.Description, dto.Options, id}

	row := d.db.QueryRowContext(ctx, query, params...)

	var updatedId int

	if err := row.Scan(&updatedId); err != nil {
		return nil, err
	}

	return &updatedId, nil
}
