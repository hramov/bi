package data_source_entity

import "time"

type Driver struct {
	Id          int
	Title       string
	Code        string
	DateCreated time.Time
	DbNeed      bool
}
