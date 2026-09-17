package travel

import (
	"testing"
	"time"
)

func TestUseCase_CreateTravel(t *testing.T) {
	repository := &MockRepository{
		CreateTravelFunc: func(travel *Travel) error {
			travel.ID = 1 // Simulate auto-increment ID
			return nil
		},
	}

	useCase := NewUseCase(repository)

	userID := uint(1)
	startDate := time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2026, 10, 21, 0, 0, 0, 0, time.UTC)

	input := CreateTravelInput{
		Title:     "東南アジア3週間",
		StartDate: startDate,
		EndDate:   endDate,
		IsPublic:  true,
	}

	travel, err := useCase.CreateTravel(userID, input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if travel.ID != 1 {
		t.Fatalf("expected ID 1, got %d", travel.ID)
	}

	if travel.UserID != userID {
		t.Fatalf("expected user ID 1, got %d", travel.UserID)
	}

	if travel.Title != input.Title {
		t.Fatalf("expected title %s, got %s", input.Title, travel.Title)
	}

	if !travel.StartDate.Equal(input.StartDate) {
		t.Fatalf("unexpected start date")
	}

	if !travel.EndDate.Equal(input.EndDate) {
		t.Fatalf("unexpected end date")
	}

	if travel.IsPublic != input.IsPublic {
		t.Fatalf("expected IsPublic true, got %v", travel.IsPublic)
	}
}
