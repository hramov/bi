package dashboard_repo

import (
	"context"
	dashboard_model "github.com/hramov/gvc-bi/backend/dashboard/internal/adapter/postgresrepo/dashboard/model"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/entity"
)

func (d *RepositoryImpl) Get(ctx context.Context) ([]*dashboard_entity.Dashboard, error) {
	query := `select * from dashboards`

	rows, err := d.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	dashboards := []*dashboard_entity.Dashboard{}

	for rows.Next() {
		m := new(dashboard_model.Model)
		err = rows.Scan(&m.Id, &m.DashId, &m.Title, &m.Description, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
		if err != nil {
			return nil, err
		}

		dash := &dashboard_entity.Dashboard{
			Id:          m.Id,
			DashId:      m.DashId,
			Title:       m.Title,
			Description: m.Description.String,
			CreatedAt:   m.CreatedAt,
			Items:       nil,
		}

		if m.UpdatedAt.Valid {
			dash.UpdatedAt = m.UpdatedAt.Time
		}

		if m.DeletedAt.Valid {
			dash.DeletedAt = m.DeletedAt.Time
		}

		dashboards = append(dashboards, dash)
	}

	return dashboards, nil
}
