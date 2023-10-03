package data_source_dto_in

import "time"

type Datasource struct {
	Id          int
	Driver      string
	DriverId    int
	Title       string
	Dsn         string
	Checked     bool
	DateCreated time.Time
}
