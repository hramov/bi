package dashboard

import (
	"database/sql"
	"time"
)

type Model struct {
	Id          int            `json:"id"`
	DashId      string         `json:"dash_id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	DeletedAt   sql.NullTime   `json:"deleted_at"`

	Items []*ItemModel `json:"items"`
}

type ItemModel struct {
	Id          int            `json:"id"`
	DashId      string         `json:"dash_id"`
	ItemType    int            `json:"type"`
	Position    sql.NullString `json:"position"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	DataQueries any            `json:"data_queries"`
	Options     any            `json:"raw_options"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	DeletedAt   sql.NullTime   `json:"deleted_at"`
}

type ItemTypeModel struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

type FormatFunctionModel struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}
