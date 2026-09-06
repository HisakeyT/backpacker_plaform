package user

import "errors"

var ErrUserNotFound = errors.New("user not found")

type Repository interface {
	Create(user *User) error
	FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(user *User) error
}
