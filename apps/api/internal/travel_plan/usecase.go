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

func (uc *UseCase) CreateTravelPlan(userID, travelID uint, input CreateTravelPlanInput) (*TravelPlan, error) {
	travelData, err := uc.travelRepository.FindByID(travelID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, travel.ErrTravelNotFound
		}
		return nil, err
	}

	if travelData.UserID != userID {
		return nil, user.ErrUserNotAuthorized
	}
	travelPlan := &TravelPlan{
		TravelID:  travelID,
		Date:      input.Date,
		Place:     input.Place,
		Content:   input.Content,
		SortOrder: input.SortOrder,
	}

	if uc.travelPlanRepository.Create(travelPlan); err != nil {
		return nil, err
	}

	return travelPlan, nil
}

func (uc *UseCase) GetTravelPlans(userID, travelID uint) ([]TravelPlan, error) {
	travelData, err := uc.travelRepository.FindByID(travelID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, travel.ErrTravelNotFound
		}
		return nil, err
	}

	if travelData.UserID != userID {
		return nil, user.ErrUserNotAuthorized
	}

	travelPlans, err := uc.travelPlanRepository.FindByTravelID(travelID)
	if err != nil {
		return nil, err
	}

	return travelPlans, nil
}

type UpdateTravelPlanInput struct {
	Date      *time.Time
	Place     *string
	Content   *string
	SortOrder *int
}

func (uc *UseCase) UpdateTravelPlan(userID, travelID, travelPlanID uint, input UpdateTravelPlanInput) (*TravelPlan, error) {
	travelData, err := uc.travelRepository.FindByID(travelID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, travel.ErrTravelNotFound
		}
		return nil, err
	}

	if travelData.UserID != userID {
		return nil, user.ErrUserNotAuthorized
	}

	travelPlan, err := uc.travelPlanRepository.FindByID(travelPlanID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTravelPlanNotFound
		}

		return nil, err
	}

	if travelPlan.TravelID != travelID {
		return nil, ErrTravelPlanNotBelongToTravel
	}

	applyTravelPlanUpdates(travelPlan, input)

	if err := uc.travelPlanRepository.Update(travelPlan); err != nil {
		return nil, err
	}

	return travelPlan, nil
}

func applyTravelPlanUpdates(travelPlan *TravelPlan, input UpdateTravelPlanInput) {
	if input.Date != nil {
		travelPlan.Date = *input.Date
	}

	if input.Place != nil {
		travelPlan.Place = *input.Place
	}

	if input.Content != nil {
		travelPlan.Content = *input.Content
	}

	if input.SortOrder != nil {
		travelPlan.SortOrder = *input.SortOrder
	}
}
