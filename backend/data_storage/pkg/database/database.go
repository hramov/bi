package database

import (
	"database/sql"
	"fmt"
	"github.com/hramov/gvc-bi/backend/datastorage/pkg/database/postgres"
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

func NewStorageForQuery(driver, dsn string) (*sql.DB, error) {
	switch driver {
	case string(Postgres):
		return postgres.New(nil, dsn)
	}
	return nil, fmt.Errorf("cannot find driver")
}
