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

func (r *GormRepository) FindByUserID(userID uint) ([]*Travel, error) {
	var travels []*Travel

	if err := r.db.Where("user_id = ?", userID).Find(&travels).Error; err != nil {
		return nil, err
	}

	return travels, nil
}

func (r *GormRepository) Update(travel *Travel) error {
	return r.db.Save(travel).Error
}

func (r *GormRepository) FindPublicTravels() ([]*Travel, error) {
	var travels []*Travel

	if err := r.db.Where("is_public = ?", true).Find(&travels).Error; err != nil {
		return nil, err
	}

	return travels, nil
}
