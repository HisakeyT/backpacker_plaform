package travel

import (
	"errors"
	"time"
)

var ErrInvalidDateRange = errors.New("start date must not be after end date")

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
	if err := validateCreateTravelInput(input); err != nil {
		return nil, err
	}

	travel := &Travel{
		UserID:    userID,
		Title:     input.Title,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		IsPublic:  input.IsPublic,
	}

	if err := u.repository.Create(travel); err != nil {
		return nil, err
	}

	return travel, nil

}

func validateCreateTravelInput(input CreateTravelInput) error {
	if input.StartDate.After(input.EndDate) {
		return ErrInvalidDateRange
	}

	return nil
}

func (u *UseCase) GetTravels(userID uint) ([]*Travel, error) {
	travels, err := u.repository.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	return travels, nil
}
