package dashboard

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	dashboard_entity "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/entity"
	dashboard_repomocks "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/mocks/repomocks"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/logger"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func randomTimestamp() time.Time {
	return time.Unix(rand.Int63n(time.Now().Unix()-94608000)+94608000, 0)
}

func generateItems() []*dashboard_entity.Item {
	var items []*dashboard_entity.Item

	for i := 0; i < 1; i++ {
		item := &dashboard_entity.Item{
			Id:          int(gofakeit.Uint8()),
			DashId:      gofakeit.UUID(),
			ItemType:    gofakeit.RandomInt([]int{0, 1}),
			Position:    gofakeit.BookTitle(),
			Title:       gofakeit.BookTitle(),
			Description: gofakeit.BookTitle(),
			DataQueries: nil,
			Options:     nil,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeletedAt:   time.Time{},
		}

		items = append(items, item)
	}

	return items

}

func generateDashboards() []*dashboard_entity.Dashboard {
	var dash []*dashboard_entity.Dashboard

	for i := 0; i < 1; i++ {
		mockDashboard := &dashboard_entity.Dashboard{
			Id:          int(gofakeit.Uint8()),
			DashId:      gofakeit.UUID(),
			Title:       gofakeit.BookTitle(),
			Description: gofakeit.LoremIpsumWord(),
			CreatedAt:   randomTimestamp(),
			UpdatedAt:   randomTimestamp(),
			DeletedAt:   randomTimestamp(),
			Items:       generateItems(),
		}

		dash = append(dash, mockDashboard)
	}

	return dash
}

func TestService_Get(t *testing.T) {
	// Create a mock repository
	repo := &dashboard_repomocks.Repository{}

	dashboards := generateDashboards()
	// Set up the expected behavior of the mock repository
	repo.On("Get", context.Background()).
		Return(dashboards, nil).
		Once()

	// Create a logger for testing
	l := logger.NewTest()

	// Create an instance of the service with the mock repository and logger
	service := NewService(repo, l)

	// Call the Get method of the service
	d, err := service.Get(context.Background())

	// Assert that the returned dashboards match the expected result
	assert.Nil(t, err)
	assert.EqualValues(t, d, dashboards)

	// Assert that the mock repository's Get method was called with the correct arguments
	repo.AssertCalled(t, "Get", context.Background())
}

func TestService_GetByDashId(t *testing.T) {
	// Create a mock repository
	repo := &dashboard_repomocks.Repository{}

	// Create a mock dashboard
	mockDashboard := &dashboard_entity.Dashboard{
		Id:          1,
		DashId:      "abc123",
		Title:       "Test Dashboard",
		Description: "This is a test dashboard",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   time.Time{},
		Items:       nil,
	}

	// Set up the expected behavior of the mock repository
	repo.On("GetByDashId", context.Background(), "abc123").
		Return(mockDashboard, nil).
		Once()

	// Create a logger for testing
	l := logger.NewTest()

	// Create an instance of the service with the mock repository and logger
	service := NewService(repo, l)

	// Call the GetByDashId method of the service
	dashboard, err := service.GetByDashId(context.Background(), "abc123")

	// Assert that the returned dashboard matches the expected result
	assert.Nil(t, err)
	assert.Equal(t, mockDashboard, dashboard)

	// Assert that the mock repository's GetByDashId method was called with the correct arguments
	repo.AssertCalled(t, "GetByDashId", context.Background(), "abc123")
}

func TestService_GetByDashId_Error(t *testing.T) {
	// Create a mock repository
	repo := &dashboard_repomocks.Repository{}

	// Set up the expected behavior of the mock repository to return an error
	expectedError := errors.New("database error")
	repo.On("GetByDashId", context.Background(), "abc123").
		Return(nil, expectedError).
		Once()

	// Create a logger for testing
	l := logger.NewTest()

	// Create an instance of the service with the mock repository and logger
	service := NewService(repo, l)

	// Call the GetByDashId method of the service
	dashboard, err := service.GetByDashId(context.Background(), "abc123")

	// Assert that the returned dashboard is nil and the error matches the expected error
	assert.Nil(t, dashboard)
	assert.Equal(t, expectedError, err)

	// Assert that the mock repository's GetByDashId method was called with the correct arguments
	repo.AssertCalled(t, "GetByDashId", context.Background(), "abc123")
}
