package database

import (
	"database/sql"
	"fmt"
	"github.com/hramov/gvc-bi/backend/datasource/pkg/database/postgres"
	"strconv"
)

type DataStorageOptions struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Sslmode  string
}

type DataSource string

const (
	Postgres DataSource = "pg"
)

func NewStorage(storage DataStorageOptions) (*sql.DB, error) {
	switch storage.Driver {
	case string(Postgres):
		port, _ := strconv.Atoi(storage.Port)

		pgOptions := &postgres.Options{
			Host:     storage.Host,
			Port:     port,
			User:     storage.User,
			Password: storage.Password,
			Database: storage.Database,
			SslMode:  storage.Sslmode,
		}

		return postgres.New(pgOptions, "")
	}
	return nil, fmt.Errorf("cannot find driver")
}

func NewStorageForQuery(storage DataStorageOptions) (*sql.DB, error) {
	switch storage.Driver {
	case string(Postgres):
		port, _ := strconv.Atoi(storage.Port)

		pgOptions := &postgres.Options{
			Host:     storage.Host,
			Port:     port,
			User:     storage.User,
			Password: storage.Password,
			Database: storage.Database,
			SslMode:  storage.Sslmode,
		}

		return postgres.NewQuery(pgOptions, "")
	}
	return nil, fmt.Errorf("cannot find driver")
}
