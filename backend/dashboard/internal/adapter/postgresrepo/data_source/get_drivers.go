package data_source_repo

import (
	"context"
	data_source_entity "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/entity"
)

func (r *RepositoryImpl) GetDrivers() ([]*data_source_entity.Driver, error) {
	query := `select * from drivers`

	rows, err := r.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}

	models := []*data_source_entity.Driver{}

	for rows.Next() {
		model := &data_source_entity.Driver{}
		err = rows.Scan(&model.Id, &model.Title, &model.Code, &model.DateCreated, &model.DbNeed)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}
