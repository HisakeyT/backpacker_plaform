package user

import (
	"errors"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestLogin_Success(t *testing.T) {
	email := "taro@example.com"

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte("password123"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		t.Fatalf("failed to generate password hash: %v", err)
	}

	mockRepo := &MockRepository{
		FindByEmailFunc: func(email string) (*User, error) {
			return &User{
				ID:           1,
				Nickname:     "Taro",
				Email:        &email,
				PasswordHash: string(hashedPassword),
			}, nil
		},
	}

	useCase := NewUseCase(mockRepo)

	user, err := useCase.Login(LoginInput{
		Email:    email,
		Password: "password123",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("expected user ID 1, got %d", user.ID)
	}

	if user.Email == nil || *user.Email != email {
		t.Errorf("expected email %s, got %v", email, user.Email)
	}
}

func TestLogin_UserNotFound(t *testing.T) {
	mockRepo := &MockRepository{
		FindByEmailFunc: func(email string) (*User, error) {
			return nil, ErrUserNotFound
		},
	}

	useCase := NewUseCase(mockRepo)

	_, err := useCase.Login(LoginInput{
		Email:    "notfound@example.com",
		Password: "password123",
	})

	if !errors.Is(err, ErrInvalidCredentials) {
		t.Fatalf("expected ErrInvalidCredentials, got %v", err)
	}
}

func TestLogin_WrongPassword(t *testing.T) {
	email := "taro@example.com"

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte("password123"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		t.Fatalf("failed to generate password hash: %v", err)
	}

	mockRepo := &MockRepository{
		FindByEmailFunc: func(email string) (*User, error) {
			return &User{
				ID:           1,
				Nickname:     "Taro",
				Email:        &email,
				PasswordHash: string(hashedPassword),
			}, nil
		},
	}

	useCase := NewUseCase(mockRepo)

	_, err = useCase.Login(LoginInput{
		Email:    email,
		Password: "wrong-password",
	})

	if !errors.Is(err, ErrInvalidCredentials) {
		t.Fatalf("expected ErrInvalidCredentials, got %v", err)
	}
}
