package travel_plan

import (
	"github.com/HisakeyT/backpacker-platform/internal/travel"
)

type Repository interface {
	Create(travelPlan *travel.TravelPlan) error
	Update(travelPlan *travel.TravelPlan) error
	Delete(id uint) error
	FindByID(id uint) (*travel.TravelPlan, error)
	FindByTravelID(travelID uint) ([]travel.TravelPlan, error)
}
