package data_source_dto_in

import "time"

type Driver struct {
	Id          int
	Title       string
	Code        string
	DateCreated time.Time
	DbNeed      bool
}
