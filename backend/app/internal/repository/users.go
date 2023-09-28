package repository

import "database/sql"

type UsersRepository struct {
	Db *sql.DB
}
