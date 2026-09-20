package travel

type MockRepository struct {
	CreateFunc   func(travel *Travel) error
	FindByIDFunc func(id uint) (*Travel, error)
}

func (m *MockRepository) Create(travel *Travel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(travel)
	}
	return nil
}

func (m *MockRepository) FindByID(id uint) (*Travel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}

	return nil, nil
}
