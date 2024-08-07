package data_source_repo

import (
	"context"
	data_source_entity "github.com/hramov/gvc-bi/backend/datasource/internal/domain/data_source/entity"
)

func (r *RepositoryImpl) Get(ctx context.Context) ([]*data_source_entity.Datasource, error) {
	query := `
		select ds.id, ds.driver_id, d.code, ds.title, ds.dsn, ds.checked, ds.date_created, d.code  
		from data_sources ds
		join drivers d on d.id = ds.driver_id
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	models := []*data_source_entity.Datasource{}

	for rows.Next() {
		model := &data_source_entity.Datasource{}
		err = rows.Scan(&model.Id, &model.DriverId, &model.DriverCode, &model.Title, &model.Dsn, &model.Checked, &model.DateCreated, &model.Driver)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}
