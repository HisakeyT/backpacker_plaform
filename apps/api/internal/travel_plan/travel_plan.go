package travel_plan

import "time"

type TravelPlan struct {
	ID        uint      `gorm:"primaryKey"`
	TravelID  uint      `gorm:"not null;index"`
	Date      time.Time `gorm:"not null"`
	Place     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	SortOrder int       `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
