package travel

type Repository interface {
	Create(travel *Travel) error
	FindByID(id uint) (*Travel, error)
	FindByUserID(userID uint) ([]*Travel, error)
}
