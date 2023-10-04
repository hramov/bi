package dashboard_entity

import "time"

type Item struct {
	Id          int       `json:"id"`
	DashId      string    `json:"dash_id"`
	ItemType    int       `json:"item_type"`
	Position    string    `json:"position"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DataQueries any       `json:"data_queries"`
	Options     any       `json:"options"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
