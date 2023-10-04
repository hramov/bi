package data_source_repo

import (
	"context"
	"database/sql"
	"errors"
	data_source_entity "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/entity"
)

func (r *RepositoryImpl) GetDriverById(ctx context.Context, id int) (*data_source_entity.Driver, error) {
	query := `select * from drivers where id = $1`
	params := []any{id}

	row := r.db.QueryRowContext(ctx, query, params...)

	if row.Err() != nil {
		if !errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, nil
		}
		return nil, row.Err()
	}

	model := &data_source_entity.Driver{}
	err := row.Scan(&model.Id, &model.Title, &model.Code, &model.DateCreated, &model.DbNeed)
	if err != nil {
		return nil, err
	}

	return model, nil
}
