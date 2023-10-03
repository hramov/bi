package dashboard_repo

import (
	"context"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/adapter/postgresrepo/dashboard/model"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/entity"
)

func (d *RepositoryImpl) GetAvailableTypes() ([]*dashboard_entity.ItemType, error) {
	query := `select * from dashboard_item_types`

	rows, err := d.db.QueryContext(context.Background(), query)

	if err != nil {
		return nil, err
	}

	types := []*dashboard_entity.ItemType{}

	for rows.Next() {
		m := dashboard_model.ItemTypeModel{}
		err = rows.Scan(&m.Id, &m.Title, &m.Name)
		if err != nil {
			return nil, err
		}

		t := &dashboard_entity.ItemType{
			Id:    m.Id,
			Title: m.Title,
			Name:  m.Name,
		}

		types = append(types, t)
	}

	return types, nil
}
