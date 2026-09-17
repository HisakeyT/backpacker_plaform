package travel

type MockRepository struct {
	CreateTravelFunc func(travel *Travel) error
}

func (m *MockRepository) CreateTravel(travel *Travel) error {
	if m.CreateTravelFunc != nil {
		return m.CreateTravelFunc(travel)
	}
	return nil
}
