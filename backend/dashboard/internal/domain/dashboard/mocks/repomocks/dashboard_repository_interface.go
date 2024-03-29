// Code generated by mockery v2.34.2. DO NOT EDIT.

package dashboard_repomocks

import (
	context "context"

	dashboard_entity "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/entity"

	dashboards_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/in"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, dto
func (_m *Repository) Create(ctx context.Context, dto dashboards_dto_in.Dashboard) (*int, error) {
	ret := _m.Called(ctx, dto)

	var r0 *int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dashboards_dto_in.Dashboard) (*int, error)); ok {
		return rf(ctx, dto)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dashboards_dto_in.Dashboard) *int); ok {
		r0 = rf(ctx, dto)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dashboards_dto_in.Dashboard) error); ok {
		r1 = rf(ctx, dto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateItem provides a mock function with given fields: ctx, dto
func (_m *Repository) CreateItem(ctx context.Context, dto dashboards_dto_in.Item) (*int, error) {
	ret := _m.Called(ctx, dto)

	var r0 *int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dashboards_dto_in.Item) (*int, error)); ok {
		return rf(ctx, dto)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dashboards_dto_in.Item) *int); ok {
		r0 = rf(ctx, dto)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dashboards_dto_in.Item) error); ok {
		r1 = rf(ctx, dto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx
func (_m *Repository) Get(ctx context.Context) ([]*dashboard_entity.Dashboard, error) {
	ret := _m.Called(ctx)

	var r0 []*dashboard_entity.Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*dashboard_entity.Dashboard, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*dashboard_entity.Dashboard); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dashboard_entity.Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAvailableTypes provides a mock function with given fields: ctx
func (_m *Repository) GetAvailableTypes(ctx context.Context) ([]*dashboard_entity.ItemType, error) {
	ret := _m.Called(ctx)

	var r0 []*dashboard_entity.ItemType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*dashboard_entity.ItemType, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*dashboard_entity.ItemType); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dashboard_entity.ItemType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByDashId provides a mock function with given fields: ctx, id
func (_m *Repository) GetByDashId(ctx context.Context, id string) (*dashboard_entity.Dashboard, error) {
	ret := _m.Called(ctx, id)

	var r0 *dashboard_entity.Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*dashboard_entity.Dashboard, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *dashboard_entity.Dashboard); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dashboard_entity.Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetItemById provides a mock function with given fields: ctx, id
func (_m *Repository) GetItemById(ctx context.Context, id int) (*dashboard_entity.Item, error) {
	ret := _m.Called(ctx, id)

	var r0 *dashboard_entity.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*dashboard_entity.Item, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *dashboard_entity.Item); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dashboard_entity.Item)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, dto, id
func (_m *Repository) Update(ctx context.Context, dto dashboards_dto_in.Dashboard, id int) (*int, error) {
	ret := _m.Called(ctx, dto, id)

	var r0 *int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dashboards_dto_in.Dashboard, int) (*int, error)); ok {
		return rf(ctx, dto, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dashboards_dto_in.Dashboard, int) *int); ok {
		r0 = rf(ctx, dto, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dashboards_dto_in.Dashboard, int) error); ok {
		r1 = rf(ctx, dto, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateItem provides a mock function with given fields: ctx, dto, id
func (_m *Repository) UpdateItem(ctx context.Context, dto dashboards_dto_in.Item, id int) (*int, error) {
	ret := _m.Called(ctx, dto, id)

	var r0 *int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dashboards_dto_in.Item, int) (*int, error)); ok {
		return rf(ctx, dto, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dashboards_dto_in.Item, int) *int); ok {
		r0 = rf(ctx, dto, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dashboards_dto_in.Item, int) error); ok {
		r1 = rf(ctx, dto, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
