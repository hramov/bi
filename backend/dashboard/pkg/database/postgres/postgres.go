package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Options struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SslMode  string
}

type postgres struct {
	options *Options
	db      *sql.DB
}

func New(options *Options, dsn string) (*sql.DB, error) {
	p := &postgres{
		options: options,
	}

	var db *sql.DB
	var err error

	if dsn != "" {
		db, err = sql.Open("postgres", dsn)
	} else {
		db, err = sql.Open("postgres", p.formatDNS())
	}

	if err != nil {
		return nil, fmt.Errorf("cannot connect to postgres: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("cannot ping postgres: %v", err)
	}
	p.db = db
	return db, nil
}

func (p *postgres) formatDNS() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", p.options.User, p.options.Password, p.options.Host, p.options.Port, p.options.Database, p.options.SslMode)
}
