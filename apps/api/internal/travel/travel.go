package travel

import (
	"errors"
	"time"
)

var ErrTravelNotFound = errors.New("travel not found")

type Travel struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Title     string    `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	IsPublic  bool      `gorm:"not null;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
