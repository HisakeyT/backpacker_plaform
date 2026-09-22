package travel_plan

import (
	"errors"
	"github.com/HisakeyT/backpacker-platform/internal/travel"
	"github.com/HisakeyT/backpacker-platform/internal/user"
	"time"

	"gorm.io/gorm"
)

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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return travel.ErrTravelNotFound
		}
		return err
	}

	if travelData.UserID != userID {
		return user.ErrUserNotAuthorized
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
