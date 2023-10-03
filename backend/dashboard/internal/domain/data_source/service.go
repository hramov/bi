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
	Create(ds data_source_dto_in.Datasource) (*int, error)
	Update(ds data_source_dto_in.Datasource, id int) (*int, error)
	Delete(id int) (*int, error)
}

type Service struct {
	repo   Repository
	logger Logger
}

func NewService(repo Repository, logger Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s Service) GetDrivers() ([]*data_source_entity.Driver, error) {
	return s.repo.GetDrivers()
}

func (s Service) GetDriverById(id int) (*data_source_entity.Driver, error) {
	return s.repo.GetDriverById(id)
}

func (s Service) Get() ([]*data_source_entity.Datasource, error) {
	return s.repo.Get()
}

func (s Service) GetById(id int) (*data_source_entity.Datasource, error) {
	return s.repo.GetById(id)
}

func (s Service) Create(ds data_source_dto_in.Datasource) (*int, error) {
	res, err := s.repo.Create(ds)
	// TODO http call to data_storage service to re-retrieve data sources from database
	return res, err
}

func (s Service) Update(ds data_source_dto_in.Datasource, id int) (*int, error) {
	return s.repo.Update(ds, id)
}

func (s Service) Delete(id int) (*int, error) {
	return s.repo.Delete(id)
}
