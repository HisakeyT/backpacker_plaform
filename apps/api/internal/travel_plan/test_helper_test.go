package travel_plan

import (
	"github.com/HisakeyT/backpacker-platform/internal/travel"
)

type MockRepository struct {
	CreateFunc         func(travelPlan *travel.TravelPlan) error
	UpdateFunc         func(travelPlan *travel.TravelPlan) error
	DeleteFunc         func(id uint) error
	FindByIDFunc       func(id uint) (*travel.TravelPlan, error)
	FindByTravelIDFunc func(travelID uint) ([]travel.TravelPlan, error)
}

func (m *MockRepository) Create(travelPlan *travel.TravelPlan) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(travelPlan)
	}
	return nil
}

func (m *MockRepository) Update(travelPlan *travel.TravelPlan) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(travelPlan)
	}

	return nil
}

func (m *MockRepository) Delete(id uint) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}

	return nil
}

func (m *MockRepository) FindByID(id uint) (*travel.TravelPlan, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}

	return nil, nil
}

func (m *MockRepository) FindByTravelID(travelID uint) ([]travel.TravelPlan, error) {
	if m.FindByTravelIDFunc != nil {
		return m.FindByTravelIDFunc(travelID)
	}

	return nil, nil
}

type MockTravelRepository struct {
	FindByIDFunc func(id uint) (*travel.Travel, error)
}

func (m *MockTravelRepository) FindByID(id uint) (*travel.Travel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return nil, nil
}
