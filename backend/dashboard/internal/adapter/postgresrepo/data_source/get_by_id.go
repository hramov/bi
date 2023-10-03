package data_source_repo

import (
	"context"
	"database/sql"
	"errors"
	data_source_entity "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/entity"
)

func (r *RepositoryImpl) GetById(id int) (*data_source_entity.Datasource, error) {
	query := `
		select ds.id, ds.driver_id, ds.title, ds.dsn, ds.checked, ds.date_created, d.code  
		from data_sources ds
		join drivers d on d.id = ds.driver_id
		where ds.id = $1
	`
	params := []any{id}

	row := r.db.QueryRowContext(context.Background(), query, params...)

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
