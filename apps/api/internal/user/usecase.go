package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	repository Repository
}

func NewUseCase(repository Repository) *UseCase {
	return &UseCase{
		repository: repository,
	}
}

type RegisterInput struct {
	Nickname             string
	Email                *string
	Password             string
	PasswordConfirmation string
}

var (
	ErrPasswordMismatch   = errors.New("password and password confirmation do not match")
	ErrEmailAlreadyUsed   = errors.New("email is already used")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func (u *UseCase) Register(input RegisterInput) (*User, error) {
	if input.Password != input.PasswordConfirmation {
		return nil, ErrPasswordMismatch
	}

	if input.Email != nil {
		existingUser, err := u.repository.FindByEmail(*input.Email)
		if err == nil && existingUser != nil {
			return nil, ErrEmailAlreadyUsed
		}

		if !errors.Is(err, ErrUserNotFound) && err != nil {
			return nil, err
		}
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &User{
		Nickname:     input.Nickname,
		Email:        input.Email,
		PasswordHash: string(passwordHash),
	}

	if err := u.repository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

type LoginInput struct {
	Email    string
	Password string
}

func (u *UseCase) Login(input LoginInput) (*User, error) {
	user, err := u.repository.FindByEmail(input.Email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, ErrInvalidCredentials
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	return user, nil
}
