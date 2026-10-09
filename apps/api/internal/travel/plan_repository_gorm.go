package travel

import (
	"gorm.io/gorm"
)

type GormPlanRepository struct {
	db *gorm.DB
}

func NewGormPlanRepository(db *gorm.DB) *GormPlanRepository {
	return &GormPlanRepository{db: db}
}

func (r *GormPlanRepository) Create(travelPlan *TravelPlan) error {
	return r.db.Create(travelPlan).Error
}

func (r *GormPlanRepository) Update(travelPlan *TravelPlan) error {
	return r.db.Save(travelPlan).Error
}

func (r *GormPlanRepository) Delete(id uint) error {
	return r.db.Delete(&TravelPlan{}, id).Error
}

func (r *GormPlanRepository) FindByID(id uint) (*TravelPlan, error) {
	var travelPlan TravelPlan

	result := r.db.First(&travelPlan, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &travelPlan, nil
}

func (r *GormPlanRepository) FindByTravelID(travelID uint) ([]TravelPlan, error) {
	var travelPlans []TravelPlan

	result := r.db.Where("travel_id = ?", travelID).Find(&travelPlans)
	if result.Error != nil {
		return nil, result.Error
	}

	return travelPlans, nil
}
