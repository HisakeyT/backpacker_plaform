package travel

type MockRepository struct {
	CreateTravelFunc func(travel *Travel) error
	FindByIDFunc     func(id uint) (*Travel, error)
}

func (m *MockRepository) CreateTravel(travel *Travel) error {
	if m.CreateTravelFunc != nil {
		return m.CreateTravelFunc(travel)
	}
	return nil
}

func (m *MockRepository) FindByID(id uint) (*Travel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}

	return nil, nil
}
