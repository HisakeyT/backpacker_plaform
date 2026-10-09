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
	travelPlanRepository PlanRepository
}

func NewUseCase(
	travelRepository TravelRepository,
	travelPlanRepository PlanRepository,
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

func (uc *UseCase) CreateTravelPlan(userID, travelID uint, input CreateTravelPlanInput) (*travel.TravelPlan, error) {
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
	travelPlan := &travel.TravelPlan{
		TravelID:  travelID,
		Date:      input.Date,
		Place:     input.Place,
		Content:   input.Content,
		SortOrder: input.SortOrder,
	}

	if err := uc.travelPlanRepository.Create(travelPlan); err != nil {
		return nil, err
	}

	return travelPlan, nil
}

func (uc *UseCase) GetTravelPlans(userID, travelID uint) ([]travel.TravelPlan, error) {
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

func (uc *UseCase) DeleteTravelPlan(userID, travelID, travelPlanID uint) error {
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

	travelPlan, err := uc.travelPlanRepository.FindByID(travelPlanID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return travel.ErrTravelPlanNotFound
		}
		return err
	}

	if travelPlan.TravelID != travelID {
		return travel.ErrTravelPlanNotBelongToTravel
	}

	return uc.travelPlanRepository.Delete(travelPlanID)
}

type UpdateTravelPlanInput struct {
	Date      *time.Time
	Place     *string
	Content   *string
	SortOrder *int
}

func (uc *UseCase) UpdateTravelPlan(userID, travelID, travelPlanID uint, input UpdateTravelPlanInput) (*travel.TravelPlan, error) {
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
			return nil, travel.ErrTravelPlanNotFound
		}

		return nil, err
	}

	if travelPlan.TravelID != travelID {
		return nil, travel.ErrTravelPlanNotBelongToTravel
	}

	applyTravelPlanUpdates(travelPlan, input)

	if err := uc.travelPlanRepository.Update(travelPlan); err != nil {
		return nil, err
	}

	return travelPlan, nil
}

func applyTravelPlanUpdates(travelPlan *travel.TravelPlan, input UpdateTravelPlanInput) {
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

func (uc *UseCase) GetPulicTravelPlan(travelID uint) ([]travel.TravelPlan, error) {
	travelData, err := uc.travelRepository.FindByID(travelID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, travel.ErrTravelNotFound
		}
		return nil, err
	}

	if !travelData.IsPublic {
		return nil, travel.ErrTravelNotFound
	}

	travelPlans, err := uc.travelPlanRepository.FindByTravelID(travelID)
	if err != nil {
		return nil, err
	}

	return travelPlans, nil
}
