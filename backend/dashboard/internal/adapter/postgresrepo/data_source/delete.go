package data_source_repo

import "context"

func (r *RepositoryImpl) Delete(id int) (*int, error) {
	query := `DELETE FROM data_sources WHERE id = $1 RETURNING id`

	params := []any{id}

	row := r.db.QueryRowContext(context.Background(), query, params...)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var deletedId int
	err := row.Scan(&deletedId)
	if err != nil {
		return nil, err
	}

	return &deletedId, nil
}
