package data_source

import (
	"context"
	"fmt"
	data_source_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/dto/in"
	data_source_entity "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/entity"
	"time"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

type Repository interface {
	GetDrivers(ctx context.Context) ([]*data_source_entity.Driver, error)
	GetDriverById(ctx context.Context, id int) (*data_source_entity.Driver, error)
	Get(ctx context.Context) ([]*data_source_entity.Datasource, error)
	GetById(ctx context.Context, id int) (*data_source_entity.Datasource, error)
	Create(ctx context.Context, ds data_source_dto_in.Datasource) (*int, error)
	Update(ctx context.Context, ds data_source_dto_in.Datasource, id int) (*int, error)
	Delete(ctx context.Context, id int) (*int, error)
}

type ExternalApi interface {
	RecallDataSource(ctx context.Context, timeout time.Duration) error
}

type Service struct {
	repo   Repository
	api    ExternalApi
	logger Logger
}

func NewService(repo Repository, api ExternalApi, logger Logger) *Service {
	return &Service{
		repo:   repo,
		api:    api,
		logger: logger,
	}
}

func (s *Service) GetDrivers(ctx context.Context) ([]*data_source_entity.Driver, error) {
	return s.repo.GetDrivers(ctx)
}

func (s *Service) GetDriverById(ctx context.Context, id int) (*data_source_entity.Driver, error) {
	return s.repo.GetDriverById(ctx, id)
}

func (s *Service) Get(ctx context.Context) ([]*data_source_entity.Datasource, error) {
	return s.repo.Get(ctx)
}

func (s *Service) GetById(ctx context.Context, id int) (*data_source_entity.Datasource, error) {
	return s.repo.GetById(ctx, id)
}

func (s *Service) Create(ctx context.Context, ds data_source_dto_in.Datasource) (*int, error) {
	res, err := s.repo.Create(ctx, ds)

	if err != nil {
		return nil, err
	}

	go s.recallDataSources(ctx)

	return res, err
}

func (s *Service) Update(ctx context.Context, ds data_source_dto_in.Datasource, id int) (*int, error) {
	res, err := s.repo.Update(ctx, ds, id)

	if err != nil {
		return nil, err
	}

	go s.recallDataSources(ctx)

	return res, err
}

func (s *Service) Delete(ctx context.Context, id int) (*int, error) {
	return s.repo.Delete(ctx, id)
}

func (s *Service) recallDataSources(ctx context.Context) {
	apiErr := s.api.RecallDataSource(ctx, 10*time.Second)
	if apiErr != nil {
		s.logger.Error(fmt.Sprintf("cannot recall data source: %v", apiErr))
		return
	}
	s.logger.Info("successfully recalled data sources")
}
