package travel_plan

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HisakeyT/backpacker-platform/internal/travel"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestHandler_GetTravelPlans(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name            string
		userID          uint
		travelID        string
		setupRepository func() (*MockTravelRepository, *MockRepository)
		wantStatus      int
	}{
		{
			name:     "正常に取得できる",
			userID:   1,
			travelID: "10",
			setupRepository: func() (*MockTravelRepository, *MockRepository) {
				travelRepository := &MockTravelRepository{
					FindByIDFunc: func(id uint) (*travel.Travel, error) {
						return &travel.Travel{
							ID:     10,
							UserID: 1,
						}, nil
					},
				}

				travelPlanRepository := &MockRepository{
					FindByTravelIDFunc: func(id uint) ([]TravelPlan, error) {
						return []TravelPlan{
							{
								ID:        1,
								TravelID:  10,
								Place:     "バンコク",
								Content:   "ワット・ポーを観光",
								SortOrder: 1,
							},
							{
								ID:        2,
								TravelID:  10,
								Place:     "バンコク",
								Content:   "王宮を観光",
								SortOrder: 2,
							},
						}, nil
					},
				}

				return travelRepository, travelPlanRepository
			},
			wantStatus: http.StatusOK,
		},
		{
			name:     "TravelPlanが0件でも取得できる",
			userID:   1,
			travelID: "10",
			setupRepository: func() (*MockTravelRepository, *MockRepository) {
				travelRepository := &MockTravelRepository{
					FindByIDFunc: func(id uint) (*travel.Travel, error) {
						return &travel.Travel{
							ID:     10,
							UserID: 1,
						}, nil
					},
				}

				travelPlanRepository := &MockRepository{
					FindByTravelIDFunc: func(id uint) ([]TravelPlan, error) {
						return []TravelPlan{}, nil
					},
				}

				return travelRepository, travelPlanRepository
			},
			wantStatus: http.StatusOK,
		},
		{
			name:     "旅行が存在しない",
			userID:   1,
			travelID: "999",
			setupRepository: func() (*MockTravelRepository, *MockRepository) {
				travelRepository := &MockTravelRepository{
					FindByIDFunc: func(id uint) (*travel.Travel, error) {
						return nil, gorm.ErrRecordNotFound
					},
				}

				travelPlanRepository := &MockRepository{}

				return travelRepository, travelPlanRepository
			},
			wantStatus: http.StatusNotFound,
		},
		{
			name:     "他人の旅行は取得できない",
			userID:   1,
			travelID: "10",
			setupRepository: func() (*MockTravelRepository, *MockRepository) {
				travelRepository := &MockTravelRepository{
					FindByIDFunc: func(id uint) (*travel.Travel, error) {
						return &travel.Travel{
							ID:     10,
							UserID: 999,
						}, nil
					},
				}

				travelPlanRepository := &MockRepository{}

				return travelRepository, travelPlanRepository
			},
			wantStatus: http.StatusForbidden,
		},
		{
			name:     "travel_idが不正",
			userID:   1,
			travelID: "abc",
			setupRepository: func() (*MockTravelRepository, *MockRepository) {
				return &MockTravelRepository{}, &MockRepository{}
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			travelRepository, travelPlanRepository := tt.setupRepository()

			useCase := NewUseCase(
				travelRepository,
				travelPlanRepository,
			)

			handler := NewHandler(useCase)

			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)

			c.Params = gin.Params{
				{
					Key:   "travel_id",
					Value: tt.travelID,
				},
			}
			c.Set("userID", tt.userID)

			c.Request = httptest.NewRequest(
				http.MethodGet,
				"/travels/"+tt.travelID+"/plans",
				nil,
			)

			handler.GetTravelPlans(c)

			if recorder.Code != tt.wantStatus {
				t.Errorf(
					"status = %d, want %d",
					recorder.Code,
					tt.wantStatus,
				)
			}
		})
	}
}
