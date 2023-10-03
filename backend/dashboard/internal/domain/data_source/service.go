package data_source

import (
	data_source_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/dto/in"
	data_source_entity "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/entity"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

type Repository interface {
	GetDrivers() ([]*data_source_entity.Driver, error)
	GetDriverById(id int) (*data_source_entity.Driver, error)
	Get() ([]*data_source_entity.Datasource, error)
	GetById(id int) (*data_source_entity.Datasource, error)
	Create(driver data_source_dto_in.Datasource) (*int, error)
	Update(driver data_source_dto_in.Datasource, id int) (*int, error)
	Delete(id int) (*int, error)
}
