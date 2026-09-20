package travel_plan

type Repository interface {
	Create(travelPlan *TravelPlan) error
}
