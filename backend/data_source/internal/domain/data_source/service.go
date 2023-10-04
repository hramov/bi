package data_source

import (
	"context"
	data_source_entity "github.com/hramov/gvc-bi/backend/datasource/internal/domain/data_source/entity"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

type Repository interface {
	Get(ctx context.Context) ([]*data_source_entity.Datasource, error)
	GetByCode(ctx context.Context, code string) (*data_source_entity.Datasource, error)
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

func (s Service) Get(ctx context.Context) ([]*data_source_entity.Datasource, error) {
	return s.repo.Get(ctx)
}

func (s Service) GetByCode(ctx context.Context, code string) (*data_source_entity.Datasource, error) {
	return s.repo.GetByCode(ctx, code)
}
