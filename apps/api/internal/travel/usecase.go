package travel

import "time"

type UseCase struct {
	repository Repository
}

func NewUseCase(repository Repository) *UseCase {
	return &UseCase{
		repository: repository,
	}
}

type CreateTravelInput struct {
	Title     string
	StartDate time.Time
	EndDate   time.Time
	IsPublic  bool
}

func (u *UseCase) CreateTravel(userID uint, input CreateTravelInput) (*Travel, error) {
	travel := &Travel{
		UserID:    userID,
		Title:     input.Title,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		IsPublic:  input.IsPublic,
	}

	if err := u.repository.CreateTravel(travel); err != nil {
		return nil, err
	}

	return travel, nil

}
