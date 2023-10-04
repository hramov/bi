package data_source_dto_out

import "time"

type Datasource struct {
	Id          int       `json:"id"`
	Driver      string    `json:"driver"`
	DriverId    int       `json:"driver_id"`
	Title       string    `json:"title"`
	Dsn         string    `json:"dsn"`
	Checked     bool      `json:"checked"`
	DateCreated time.Time `json:"date_created"`
}
