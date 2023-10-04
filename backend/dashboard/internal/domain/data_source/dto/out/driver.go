package data_source_dto_out

import "time"

type Driver struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Code        string    `json:"code"`
	DateCreated time.Time `json:"date_created"`
	DbNeed      bool      `json:"db_need"`
}
