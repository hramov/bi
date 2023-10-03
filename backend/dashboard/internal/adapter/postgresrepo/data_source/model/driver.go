package data_source_model

import "time"

type Driver struct {
	Id          int
	Title       string
	Code        string
	DateCreated time.Time
	DbNeed      bool
}
