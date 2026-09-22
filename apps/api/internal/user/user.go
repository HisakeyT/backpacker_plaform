package user

import (
	"errors"
	"time"
)

var ErrUserNotAuthorized = errors.New("user does not have permission to create a travel plan for this travel")

type User struct {
	ID           uint
	Nickname     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
