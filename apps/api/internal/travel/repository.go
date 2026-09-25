package travel

type Repository interface {
	Create(travel *Travel) error
	FindByID(id uint) (*Travel, error)
	FindByUserID(userID uint) ([]*Travel, error)
	Update(travel *Travel) error
	FindPublicTravels() ([]*Travel, error)
}
