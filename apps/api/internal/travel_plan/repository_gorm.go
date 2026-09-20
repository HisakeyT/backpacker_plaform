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
