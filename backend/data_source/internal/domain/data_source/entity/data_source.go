package data_source_entity

import "time"

type Datasource struct {
	Id          int
	Driver      string
	DriverId    int
	Title       string
	PluginName  string
	Dsn         string
	Checked     bool
	DateCreated time.Time
}
