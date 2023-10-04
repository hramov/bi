package dashboard

import (
	"context"
	dashboards_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/in"
	dashboard_entity "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/entity"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

type Repository interface {
	Get(ctx context.Context) ([]*dashboard_entity.Dashboard, error)
	GetItemById(ctx context.Context, id int) (*dashboard_entity.Item, error)
	GetAvailableTypes(ctx context.Context) ([]*dashboard_entity.ItemType, error)
	GetByDashId(ctx context.Context, id string) (*dashboard_entity.Dashboard, error)
	Create(ctx context.Context, dto dashboards_dto_in.Dashboard) (*int, error)
	Update(ctx context.Context, dto dashboards_dto_in.Dashboard, id int) (*int, error)
	CreateItem(ctx context.Context, dto dashboards_dto_in.Item) (*int, error)
	UpdateItem(ctx context.Context, dto dashboards_dto_in.Item, id int) (*int, error)
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

func (s *Service) Get(ctx context.Context) ([]*dashboard_entity.Dashboard, error) {
	return s.repo.Get(ctx)
}

func (s *Service) GetByDashId(ctx context.Context, id string) (*dashboard_entity.Dashboard, error) {
	dash, err := s.repo.GetByDashId(ctx, id)
	if err != nil {
		s.logger.Error(err.Error())
	}
	return dash, err
}

func (s *Service) Create(ctx context.Context, dto dashboards_dto_in.Dashboard) (*int, error) {
	return s.repo.Create(ctx, dto)
}

func (s *Service) Update(ctx context.Context, dto dashboards_dto_in.Dashboard, id int) (*int, error) {
	return s.repo.Update(ctx, dto, id)
}

func (s *Service) GetItemById(ctx context.Context, id int) (*dashboard_entity.Item, error) {
	return s.repo.GetItemById(ctx, id)
}

func (s *Service) CreateItem(ctx context.Context, dto dashboards_dto_in.Item) (*int, error) {
	return s.repo.CreateItem(ctx, dto)
}

func (s *Service) UpdateItem(ctx context.Context, dto dashboards_dto_in.Item, id int) (*int, error) {
	return s.repo.UpdateItem(ctx, dto, id)
}

func (s *Service) GetAvailableTypes(ctx context.Context) ([]*dashboard_entity.ItemType, error) {
	return s.repo.GetAvailableTypes(ctx)
}
