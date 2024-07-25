package data_source_entity

import "time"

type Datasource struct {
	Id          int
	DriverCode  string
	Driver      string
	DriverId    int
	Title       string
	Dsn         string
	Checked     bool
	DateCreated time.Time
}
