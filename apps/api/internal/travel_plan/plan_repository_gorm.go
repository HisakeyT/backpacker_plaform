package travel_plan

import (
	"gorm.io/gorm"

	"github.com/HisakeyT/backpacker-platform/internal/travel"
)

type GormPlanRepository struct {
	db *gorm.DB
}

func NewGormPlanRepository(db *gorm.DB) *GormPlanRepository {
	return &GormPlanRepository{db: db}
}

func (r *GormPlanRepository) Create(travelPlan *travel.TravelPlan) error {
	return r.db.Create(travelPlan).Error
}

func (r *GormPlanRepository) Update(travelPlan *travel.TravelPlan) error {
	return r.db.Save(travelPlan).Error
}

func (r *GormPlanRepository) Delete(id uint) error {
	return r.db.Delete(&travel.TravelPlan{}, id).Error
}

func (r *GormPlanRepository) FindByID(id uint) (*travel.TravelPlan, error) {
	var travelPlan travel.TravelPlan

	result := r.db.First(&travelPlan, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &travelPlan, nil
}

func (r *GormPlanRepository) FindByTravelID(travelID uint) ([]travel.TravelPlan, error) {
	var travelPlans []travel.TravelPlan

	result := r.db.Where("travel_id = ?", travelID).Find(&travelPlans)
	if result.Error != nil {
		return nil, result.Error
	}

	return travelPlans, nil
}
