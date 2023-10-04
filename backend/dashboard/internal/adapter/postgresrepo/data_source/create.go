package data_source_repo

import (
	"context"
	data_source_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/dto/in"
)

func (r *RepositoryImpl) Create(ctx context.Context, ds data_source_dto_in.Datasource) (*int, error) {
	query := `
		INSERT INTO data_sources (driver_id, title, dsn, checked)
		SELECT id, $2, $3, $4 from drivers WHERE code = $1
		RETURNING id
	`

	params := []any{ds.Driver, ds.Title, ds.Dsn, ds.Checked}

	row := r.db.QueryRowContext(ctx, query, params...)
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
