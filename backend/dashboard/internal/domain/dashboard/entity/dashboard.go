package dashboard_entity

import "time"

type Dashboard struct {
	Id          int       `json:"id"`
	DashId      string    `json:"dash_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`

	Items []*Item `json:"items"`
}
