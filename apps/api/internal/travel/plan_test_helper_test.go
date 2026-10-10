package travel

type MockPlanRepository struct {
	CreateFunc         func(travelPlan *TravelPlan) error
	UpdateFunc         func(travelPlan *TravelPlan) error
	DeleteFunc         func(id uint) error
	FindByIDFunc       func(id uint) (*TravelPlan, error)
	FindByTravelIDFunc func(travelID uint) ([]TravelPlan, error)
}

func (m *MockPlanRepository) Create(travelPlan *TravelPlan) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(travelPlan)
	}
	return nil
}

func (m *MockPlanRepository) Update(travelPlan *TravelPlan) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(travelPlan)
	}

	return nil
}

func (m *MockPlanRepository) Delete(id uint) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}

	return nil
}

func (m *MockPlanRepository) FindByID(id uint) (*TravelPlan, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}

	return nil, nil
}

func (m *MockPlanRepository) FindByTravelID(travelID uint) ([]TravelPlan, error) {
	if m.FindByTravelIDFunc != nil {
		return m.FindByTravelIDFunc(travelID)
	}

	return nil, nil
}

type MockTravelRepository struct {
	FindByIDFunc func(id uint) (*Travel, error)
}

func (m *MockTravelRepository) FindByID(id uint) (*Travel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return nil, nil
}
