package dashboard_repo

import (
	"context"
	dashboards_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/in"
)

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
