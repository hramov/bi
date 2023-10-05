package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
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
	var err error = errors.New("trying to connect to database")

	counter := 0

	for err != nil {
		time.Sleep(5 * time.Second)

		if dsn != "" {
			db, err = sql.Open("postgres", dsn)
		} else {
			db, err = sql.Open("postgres", p.formatDNS())
		}

		if err != nil {
			log.Printf("cannot connect to postgres: %v\n", err)
			counter++
			if counter >= 5 {
				return nil, fmt.Errorf("cannot pind postgres")
			}
			continue
		}

		err = db.Ping()

		if err != nil {
			log.Printf("cannot ping postgres: %v\n", err)
			counter++
			if counter >= 5 {
				return nil, fmt.Errorf("cannot pind postgres")
			}
			continue
		}
	}

	p.db = db
	return db, nil
}

func NewQuery(options *Options, dsn string) (*sql.DB, error) {
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
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	p.db = db
	return db, nil
}

func (p *postgres) formatDNS() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", p.options.User, p.options.Password, p.options.Host, p.options.Port, p.options.Database, p.options.SslMode)
}
