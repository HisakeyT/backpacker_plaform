package travel

type PlanRepository interface {
	Create(travelPlan *TravelPlan) error
	Update(travelPlan *TravelPlan) error
	Delete(id uint) error
	FindByID(id uint) (*TravelPlan, error)
	FindByTravelID(travelID uint) ([]TravelPlan, error)
}
