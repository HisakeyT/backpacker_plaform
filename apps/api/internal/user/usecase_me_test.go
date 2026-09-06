package user

import (
	"testing"

	"github.com/HisakeyT/backpacker-platform/internal/auth"
)

func TestUseCase_Me(t *testing.T) {
	repository := &fakeRepository{
		user: &User{
			ID:       1,
			Nickname: "test",
			Email:    "test@example.com",
		},
	}

	jwtManager := auth.NewJWTManager("secret")
	useCase := NewUseCase(repository, jwtManager)

	user, err := useCase.Me(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.ID != 1 {
		t.Errorf("expected user ID %d, got %d", 1, user.ID)
	}
}
