package travel

import (
	"errors"
	"reflect"
	"testing"

	"github.com/HisakeyT/backpacker-platform/internal/user"
)

func TestUseCase_GetTravel(t *testing.T) {
	userID := uint(1)

	tests := []struct {
		name            string
		travelID        uint
		repository      *Travel
		repositoryError error
		want            *Travel
		wantErr         error
	}{
		{
			name:     "自分の旅行を取得できる",
			travelID: 1,
			repository: &Travel{
				ID:     1,
				UserID: userID,
				Title:  "旅行1",
			},
			want: &Travel{
				ID:     1,
				UserID: userID,
				Title:  "旅行1",
			},
		},
		{
			name:            "旅行が存在しない",
			travelID:        999,
			repositoryError: ErrTravelNotFound,
			wantErr:         ErrTravelNotFound,
		},
		{
			name:     "他人の旅行を取得しようとした",
			travelID: 2,
			repository: &Travel{
				ID:     2,
				UserID: 2,
				Title:  "他人の旅行",
			},
			wantErr: user.ErrUserNotAuthorized,
		},
		{
			name:            "Repositoryでエラーが発生した場合、UseCaseでもエラーを返す",
			travelID:        1,
			repositoryError: errors.New("database error"),
			wantErr:         errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &MockRepository{
				FindByIDFunc: func(id uint) (*Travel, error) {
					return tt.repository, tt.repositoryError
				},
			}

			useCase := NewUseCase(repository)
			travel, err := useCase.GetTravel(userID, tt.travelID)

			if tt.repositoryError != nil {
				if !errors.Is(err, tt.repositoryError) {
					t.Fatalf("expected error %v, got %v", tt.repositoryError, err)
				}

				return
			}

			if tt.want != nil {
				if !reflect.DeepEqual(travel, tt.want) {
					t.Fatalf("expected %+v, got %+v", tt.want, travel)
				}
			}
		})
	}
}
