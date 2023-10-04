package data_source_repo

import (
	"context"
	"database/sql"
	"errors"
	data_source_entity "github.com/hramov/gvc-bi/backend/datasource/internal/domain/data_source/entity"
)

func (r *RepositoryImpl) GetByCode(ctx context.Context, code string) (*data_source_entity.Datasource, error) {
	query := `
		select ds.id, ds.driver_id, ds.title, ds.dsn, ds.checked, ds.date_created, d.code  
		from data_sources ds
		join drivers d on d.id = ds.driver_id
		where d.code = $1
	`
	params := []any{code}

	row := r.db.QueryRowContext(ctx, query, params...)

	if row.Err() != nil {
		if !errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, nil
		}
		return nil, row.Err()
	}

	model := &data_source_entity.Datasource{}
	err := row.Scan(&model.Id, &model.DriverId, &model.Title, &model.Dsn, &model.Checked, &model.DateCreated, &model.Driver)
	if err != nil {
		return nil, err
	}

	return model, nil
}
