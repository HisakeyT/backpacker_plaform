package travel

import "gorm.io/gorm"

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Create(travel *Travel) error {
	return r.db.Create(travel).Error
}

func (r *GormRepository) FindByID(id uint) (*Travel, error) {
	var travel Travel

	if err := r.db.First(&travel, id).Error; err != nil {
		return nil, err
	}

	return &travel, nil
}
