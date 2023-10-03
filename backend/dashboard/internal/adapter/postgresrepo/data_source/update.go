package data_source_repo

import (
	"context"
	data_source_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/dto/in"
)

func (r *RepositoryImpl) Update(ds data_source_dto_in.Datasource, id int) (*int, error) {
	query := `
		UPDATE data_sources
		SET driver_id = (SELECT id FROM drivers WHERE code = $1), title = $2, dsn = $3, checked = $4
		WHERE id = $5
		RETURNING id
	`

	params := []any{ds.Driver, ds.Title, ds.Dsn, ds.Checked, id}

	row := r.db.QueryRowContext(context.Background(), query, params...)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var updatedId int
	err := row.Scan(&updatedId)
	if err != nil {
		return nil, err
	}

	return &updatedId, nil
}
