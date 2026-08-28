package user

type MockRepository struct {
	CreateFunc      func(user *User) error
	FindByIDFunc    func(id uint) (*User, error)
	FindByEmailFunc func(email string) (*User, error)
}

func (m *MockRepository) Create(user *User) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(user)
	}
	return nil
}

func (m *MockRepository) FindByID(id uint) (*User, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return nil, ErrUserNotFound
}

func (m *MockRepository) FindByEmail(email string) (*User, error) {
	if m.FindByEmailFunc != nil {
		return m.FindByEmailFunc(email)
	}
	return nil, ErrUserNotFound
}
