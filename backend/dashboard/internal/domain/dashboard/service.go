package dashboard

import dashboards_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/in"

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

type Repository interface {
	Get() ([]*Dashboard, error)
	GetItemById(id int) (*Item, error)
	GetAvailableTypes() ([]*ItemType, error)
	GetByDashId(id string) (*Dashboard, error)
	Create(dto dashboards_dto_in.Dashboard) (*int, error)
	Update(dto dashboards_dto_in.Dashboard, id int) (*int, error)
	CreateItem(dto dashboards_dto_in.Item) (*int, error)
	UpdateItem(dto dashboards_dto_in.Item, id int) (*int, error)
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

func (s *Service) Get() ([]*Dashboard, error) {
	return s.repo.Get()
}

func (s *Service) GetByDashId(id string) (*Dashboard, error) {
	dash, err := s.repo.GetByDashId(id)
	if err != nil {
		s.logger.Error(err.Error())
	}
	return dash, err
}

func (s *Service) Create(dto dashboards_dto_in.Dashboard) (*int, error) {
	return s.repo.Create(dto)
}

func (s *Service) Update(dto dashboards_dto_in.Dashboard, id int) (*int, error) {
	return s.repo.Update(dto, id)
}

func (s *Service) GetItemById(id int) (*Item, error) {
	return s.repo.GetItemById(id)
}

func (s *Service) CreateItem(dto dashboards_dto_in.Item) (*int, error) {
	return s.repo.CreateItem(dto)
}

func (s *Service) UpdateItem(dto dashboards_dto_in.Item, id int) (*int, error) {
	return s.repo.UpdateItem(dto, id)
}

func (s *Service) GetAvailableTypes() ([]*ItemType, error) {
	return s.repo.GetAvailableTypes()
}
