package dashboard_model

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
