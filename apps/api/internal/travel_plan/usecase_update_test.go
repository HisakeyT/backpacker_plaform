package travel_plan

import (
	"errors"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"

	"github.com/HisakeyT/backpacker-platform/internal/travel"
	"github.com/HisakeyT/backpacker-platform/internal/user"
)

func TestUseCase_UpdateTravelPlan(t *testing.T) {
	userID := uint(1)
	travelID := uint(1)
	travelPlanID := uint(1)

	date := time.Date(2026, 10, 2, 0, 0, 0, 0, time.UTC)
	newDate := time.Date(2026, 10, 3, 0, 0, 0, 0, time.UTC)

	place := "バンコク"
	newPlace := "チェンマイ"

	content := "ワット・ポー"
	newContent := "ワット・プラ・シン"

	sortOrder := 1
	newSortOrder := 2

	t.Run("正常系", func(t *testing.T) {
		tests := []struct {
			name        string
			updateInput travel.UpdateTravelPlanInput
			want        *travel.TravelPlan
		}{
			{
				name: "日付を更新する",
				updateInput: travel.UpdateTravelPlanInput{
					Date: &newDate,
				},
				want: &travel.TravelPlan{
					ID:        travelPlanID,
					TravelID:  travelID,
					Date:      newDate,
					Place:     place,
					Content:   content,
					SortOrder: sortOrder,
				},
			},
			{
				name: "場所を更新する",
				updateInput: travel.UpdateTravelPlanInput{
					Place: &newPlace,
				},
				want: &travel.TravelPlan{
					ID:        travelPlanID,
					TravelID:  travelID,
					Date:      date,
					Place:     newPlace,
					Content:   content,
					SortOrder: sortOrder,
				},
			},
			{
				name: "内容を更新する",
				updateInput: travel.UpdateTravelPlanInput{
					Content: &newContent,
				},
				want: &travel.TravelPlan{
					ID:        travelPlanID,
					TravelID:  travelID,
					Date:      date,
					Place:     place,
					Content:   newContent,
					SortOrder: sortOrder,
				},
			},
			{
				name: "順番を更新する",
				updateInput: travel.UpdateTravelPlanInput{
					SortOrder: &newSortOrder,
				},
				want: &travel.TravelPlan{
					ID:        travelPlanID,
					TravelID:  travelID,
					Date:      date,
					Place:     place,
					Content:   content,
					SortOrder: newSortOrder,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				repository := &MockRepository{
					FindByIDFunc: func(id uint) (*travel.TravelPlan, error) {
						return &travel.TravelPlan{
							ID:        travelPlanID,
							TravelID:  travelID,
							Date:      date,
							Place:     place,
							Content:   content,
							SortOrder: sortOrder,
						}, nil
					},
					UpdateFunc: func(travelPlan *travel.TravelPlan) error {
						return nil
					},
				}

				travelRepository := &MockTravelRepository{
					FindByIDFunc: func(id uint) (*travel.Travel, error) {
						return &travel.Travel{
							ID:     travelID,
							UserID: userID,
						}, nil
					},
				}

				useCase := travel.NewPlanUseCase(travelRepository, repository)

				got, err := useCase.UpdateTravelPlan(
					userID,
					travelID,
					travelPlanID,
					tt.updateInput,
				)

				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("expected %+v, got %+v", tt.want, got)
				}
			})
		}
	})

	t.Run("異常系", func(t *testing.T) {
		databaseError := errors.New("database error")
		updateError := errors.New("update error")

		tests := []struct {
			name                    string
			travelRepository        *travel.Travel
			travelRepositoryErr     error
			travelPlanRepository    *travel.TravelPlan
			travelPlanRepositoryErr error
			updateErr               error
			wantErr                 error
		}{
			{
				name:                "Travelが存在しない",
				travelRepositoryErr: gorm.ErrRecordNotFound,
				wantErr:             travel.ErrTravelNotFound,
			},
			{
				name: "他人のTravelを更新しようとする",
				travelRepository: &travel.Travel{
					ID:     travelID,
					UserID: 2,
				},
				wantErr: user.ErrUserNotAuthorized,
			},
			{
				name:                "TravelのRepositoryでエラーが発生する",
				travelRepositoryErr: databaseError,
				wantErr:             databaseError,
			},
			{
				name: "TravelPlanが存在しない",
				travelRepository: &travel.Travel{
					ID:     travelID,
					UserID: userID,
				},
				travelPlanRepositoryErr: gorm.ErrRecordNotFound,
				wantErr:                 travel.ErrTravelPlanNotFound,
			},
			{
				name: "TravelPlanのRepositoryでエラーが発生する",
				travelRepository: &travel.Travel{
					ID:     travelID,
					UserID: userID,
				},
				travelPlanRepositoryErr: databaseError,
				wantErr:                 databaseError,
			},
			{
				name: "TravelPlanが指定されたTravelに属していない",
				travelRepository: &travel.Travel{
					ID:     travelID,
					UserID: userID,
				},
				travelPlanRepository: &travel.TravelPlan{
					ID:       travelPlanID,
					TravelID: 2,
				},
				wantErr: travel.ErrTravelPlanNotBelongToTravel,
			},
			{
				name: "TravelPlanのUpdateでエラーが発生する",
				travelRepository: &travel.Travel{
					ID:     travelID,
					UserID: userID,
				},
				travelPlanRepository: &travel.TravelPlan{
					ID:       travelPlanID,
					TravelID: travelID,
				},
				updateErr: updateError,
				wantErr:   updateError,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				repository := &MockRepository{
					FindByIDFunc: func(id uint) (*travel.TravelPlan, error) {
						return tt.travelPlanRepository, tt.travelPlanRepositoryErr
					},
					UpdateFunc: func(travelPlan *travel.TravelPlan) error {
						return tt.updateErr
					},
				}

				travelRepository := &MockTravelRepository{
					FindByIDFunc: func(id uint) (*travel.Travel, error) {
						return tt.travelRepository, tt.travelRepositoryErr
					},
				}

				useCase := travel.NewPlanUseCase(travelRepository, repository)

				_, err := useCase.UpdateTravelPlan(
					userID,
					travelID,
					travelPlanID,
					travel.UpdateTravelPlanInput{
						Place: &newPlace,
					},
				)

				if !errors.Is(err, tt.wantErr) {
					t.Fatalf(
						"expected error %v, got %v",
						tt.wantErr,
						err,
					)
				}
			})
		}
	})
}
