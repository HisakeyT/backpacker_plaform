package user

import "time"

type User struct {
	ID           uint
	Nickname     string
	Email        *string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
