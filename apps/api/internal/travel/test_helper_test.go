package travel

type MockRepository struct {
	CreateFunc       func(travel *Travel) error
	FindByIDFunc     func(id uint) (*Travel, error)
	FindByUserIDFunc func(userID uint) ([]*Travel, error)
	UpdateFunc       func(travel *Travel) error
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

func (m *MockRepository) FindByUserID(userID uint) ([]*Travel, error) {
	if m.FindByUserIDFunc != nil {
		return m.FindByUserIDFunc(userID)
	}

	return nil, nil
}

func (m *MockRepository) Update(travel *Travel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(travel)
	}

	return nil
}
