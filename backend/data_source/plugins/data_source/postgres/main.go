package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/hramov/gvc-bi/backend/datasource/internal/domain/data_source/connections"
	_ "github.com/lib/pq"
)

type Postgres struct {
}

func New() connections.DataSourcePlugin {
	return &Postgres{}
}

func (p *Postgres) Connect(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to postgres (%s): %v", dsn, err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot ping postgres (%s): %v", dsn, err)
	}

	return db, nil
}
