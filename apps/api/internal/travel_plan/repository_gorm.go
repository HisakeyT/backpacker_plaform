package travel_plan

import "gorm.io/gorm"

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Create(travelPlan *TravelPlan) error {
	return r.db.Create(travelPlan).Error
}

func (r *GormRepository) Update(travelPlan *TravelPlan) error {
	return r.db.Save(travelPlan).Error
}

func (r *GormRepository) FindByID(id uint) (*TravelPlan, error) {
	var travelPlan TravelPlan

	result := r.db.First(&travelPlan, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &travelPlan, nil
}

func (r *GormRepository) FindByTravelID(travelID uint) ([]TravelPlan, error) {
	var travelPlans []TravelPlan

	result := r.db.Where("travel_id = ?", travelID).Find(&travelPlans)
	if result.Error != nil {
		return nil, result.Error
	}

	return travelPlans, nil
}
