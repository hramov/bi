package storage

import (
	"database/sql"
	"fmt"
	"github.com/hramov/gvc-bi/ds/internal/storage/postgres"
)

var storages map[string]*sql.DB = make(map[string]*sql.DB)

type SourceType string

const (
	Postgres SourceType = "postgres"
	Mssql    SourceType = "mssql"
)

func New(sourceName, sourceType string, options postgres.Options) (*sql.DB, error) {
	if source, ok := storages[sourceName]; ok {
		return source, nil
	}

	switch sourceType {
	case string(Postgres):
		source, err := postgres.New(options)
		if err != nil {
			return nil, fmt.Errorf("cannot create postgres connection: %v", err)
		}
		storages[string(Postgres)] = source
		return source, nil
	}

	return nil, fmt.Errorf("data source type not found")
}
