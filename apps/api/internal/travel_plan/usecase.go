package travel_plan

import (
	"errors"
	"github.com/HisakeyT/backpacker-platform/internal/travel"
	"time"
)

var ErrUserNotAuthorized = errors.New("user does not have permission to create a travel plan for this travel")

type TravelRepository interface {
	FindByID(id uint) (*travel.Travel, error)
}

type UseCase struct {
	travelRepository     TravelRepository
	travelPlanRepository Repository
}

func NewUseCase(
	travelRepository TravelRepository,
	travelPlanRepository Repository,
) *UseCase {
	return &UseCase{
		travelRepository:     travelRepository,
		travelPlanRepository: travelPlanRepository,
	}
}

type CreateTravelPlanInput struct {
	Date      time.Time
	Place     string
	Content   string
	SortOrder int
}

func (uc *UseCase) CreateTravelPlan(userID, travelID uint, input CreateTravelPlanInput) error {
	travelData, err := uc.travelRepository.FindByID(travelID)
	if err != nil {
		return err
	}

	if travelData.UserID != userID {
		return ErrUserNotAuthorized
	}
	travelPlan := &TravelPlan{
		TravelID:  travelID,
		Date:      input.Date,
		Place:     input.Place,
		Content:   input.Content,
		SortOrder: input.SortOrder,
	}

	return uc.travelPlanRepository.Create(travelPlan)
}
