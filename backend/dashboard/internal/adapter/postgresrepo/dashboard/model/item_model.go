package dashboard_model

import (
	"database/sql"
	"time"
)

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
