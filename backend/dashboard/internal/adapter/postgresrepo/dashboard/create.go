package dashboard_repo

import (
	"context"
	dashboards_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/in"
)

func (d *RepositoryImpl) Create(ctx context.Context, dto dashboards_dto_in.Dashboard) (*int, error) {
	query := `
		insert into dashboards (title, description)
		values ($1, $2)
		returning id;
	`
	params := []any{dto.Title, dto.Description}

	row := d.db.QueryRowContext(ctx, query, params...)

	var id int
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}
