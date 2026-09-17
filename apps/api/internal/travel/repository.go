package travel

type Repository interface {
	CreateTravel(travel *Travel) error
}
