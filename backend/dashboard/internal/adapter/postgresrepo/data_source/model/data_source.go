package data_source_model

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
