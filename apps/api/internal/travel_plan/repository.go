package travel_plan

type Repository interface {
	Create(travelPlan *TravelPlan) error
	Update(travelPlan *TravelPlan) error
	Delete(id uint) error
	FindByID(id uint) (*TravelPlan, error)
	FindByTravelID(travelID uint) ([]TravelPlan, error)
}
