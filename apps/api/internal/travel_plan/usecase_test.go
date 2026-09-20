package travel_plan

import (
	"errors"
	"testing"
	"time"

	"github.com/HisakeyT/backpacker-platform/internal/travel"
)

var errRepository = errors.New("repository error")

func TestUseCase_CreateTravelPlan(t *testing.T) {
	t.Run("正常に作成できる", func(t *testing.T) {
		// Arrange
		userID := uint(1)
		travelID := uint(10)

		input := CreateTravelPlanInput{
			Date:      time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC),
			Place:     "バンコク",
			Content:   "ワット・ポーを観光",
			SortOrder: 1,
		}

		travelData := &travel.Travel{
			ID:     travelID,
			UserID: userID,
		}

		travelRepository := &MockTravelRepository{
			FindByIDFunc: func(id uint) (*travel.Travel, error) {
				return travelData, nil
			},
		}

		var got *TravelPlan

		travelPlanRepository := &MockRepository{
			CreateFunc: func(travelPlan *TravelPlan) error {
				got = travelPlan
				return nil
			},
		}

		useCase := NewUseCase(
			travelRepository,
			travelPlanRepository,
		)

		// Act
		err := useCase.CreateTravelPlan(userID, travelID, input)

		// Assert
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got.TravelID != travelID {
			t.Errorf("TravelID = %d, want %d", got.TravelID, travelID)
		}

		if !got.Date.Equal(input.Date) {
			t.Errorf("Date = %v, want %v", got.Date, input.Date)
		}

		if got.Place != input.Place {
			t.Errorf("Place = %s, want %s", got.Place, input.Place)
		}

		if got.Content != input.Content {
			t.Errorf("Content = %s, want %s", got.Content, input.Content)
		}

		if got.SortOrder != input.SortOrder {
			t.Errorf("SortOrder = %d, want %d", got.SortOrder, input.SortOrder)
		}
	})
}

func TestUseCase_CreateTravelPlan_Invalid(t *testing.T) {
	tests := []struct {
		name        string
		userID      uint
		travelID    uint
		travelData  *travel.Travel
		findByIDErr error
		wantErr     error
	}{
		{
			name:     "他人の旅行には旅程を追加できない",
			userID:   1,
			travelID: 10,
			travelData: &travel.Travel{
				ID:     10,
				UserID: 2,
			},
			wantErr: ErrUserNotAuthorized,
		},
		{
			name:        "旅行の取得に失敗した場合",
			userID:      1,
			travelID:    10,
			findByIDErr: errRepository,
			wantErr:     errRepository,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			travelRepository := &MockTravelRepository{
				FindByIDFunc: func(id uint) (*travel.Travel, error) {
					return tt.travelData, tt.findByIDErr
				},
			}

			travelPlanRepository := &MockRepository{
				CreateFunc: func(travelPlan *TravelPlan) error {
					t.Fatal("Create should not be called")
					return nil
				},
			}

			useCase := NewUseCase(
				travelRepository,
				travelPlanRepository,
			)

			input := CreateTravelPlanInput{
				Date:      time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC),
				Place:     "バンコク",
				Content:   "観光",
				SortOrder: 1,
			}

			err := useCase.CreateTravelPlan(
				tt.userID,
				tt.travelID,
				input,
			)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
