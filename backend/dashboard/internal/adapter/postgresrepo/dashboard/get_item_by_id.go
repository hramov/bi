package dashboard_repo

import (
	"context"
	"database/sql"
	"errors"
	dashboard_entity "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/entity"
)

func (d *RepositoryImpl) GetItemById(id int) (*dashboard_entity.Item, error) {
	query := `select * from dashboard_items where id = $1`
	params := []any{id}

	row := d.db.QueryRowContext(context.Background(), query, params...)

	if row.Err() != nil {
		if !errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, nil
		}
		return nil, row.Err()
	}

	model := new(dashboard_entity.Item)
	err := row.Scan(&model.DashId, &model.ItemType, &model.Position, &model.Title, &model.Description, &model.Options, &model.CreatedAt, &model.UpdatedAt, &model.DeletedAt, &model.Id)
	if err != nil {
		return nil, err
	}

	return model, nil
}
