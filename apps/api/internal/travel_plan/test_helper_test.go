package travel_plan

import (
	"github.com/HisakeyT/backpacker-platform/internal/travel"
)

type MockRepository struct {
	CreateFunc         func(travelPlan *TravelPlan) error
	FindByTravelIDFunc func(travelID uint) ([]TravelPlan, error)
}

func (m *MockRepository) Create(travelPlan *TravelPlan) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(travelPlan)
	}
	return nil
}

func (m *MockRepository) FindByTravelID(travelID uint) ([]TravelPlan, error) {
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
