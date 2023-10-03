package dashboard

import (
	"time"
)

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

type ItemType struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

type FormatFunction struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}
