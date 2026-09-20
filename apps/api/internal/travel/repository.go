package travel

type Repository interface {
	CreateTravel(travel *Travel) error
	FindByID(id uint) (*Travel, error)
}
