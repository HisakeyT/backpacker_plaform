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

func (r *GormRepository) Delete(id uint) error {
	if err := r.db.Delete(&Travel{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *GormRepository) FindPublicTravels() ([]*PublicTravel, error) {
	var PublicTravels []*PublicTravel

	if err := r.db.Model(&Travel{}).
		Select("travels.*, users.nickname AS author_nickname").
		Joins("JOIN users ON users.id = travels.user_id").
		Where("travels.is_public = ?", true).
		Order("travels.created_at DESC").
		Scan(&PublicTravels).Error; err != nil {
		return nil, err
	}

	return PublicTravels, nil
}

func (r *GormRepository) FindPublicByID(id uint) (*PublicTravel, error) {
	var publicTravel PublicTravel

	result := r.db.Model(&Travel{}).
		Select("travels.*, users.nickname AS author_nickname").
		Joins("JOIN users ON users.id = travels.user_id").
		Where("travels.is_public = ? AND travels.id = ?", true, id).
		Limit(1).
		Scan(&publicTravel)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &publicTravel, nil
}
