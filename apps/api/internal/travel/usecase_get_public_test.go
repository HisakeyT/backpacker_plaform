package travel

import (
	"errors"
	"reflect"
	"testing"
)

func TestUseCase_GetPublicTravels(t *testing.T) {
	tests := []struct {
		name        string
		repository  Repository
		wantTravels []*Travel
		wantErr     bool
	}{
		{
			name: "公開Travelを複数取得できる",
			repository: &MockRepository{
				FindPublicTravelsFunc: func() ([]*Travel, error) {
					return []*Travel{
						{
							ID:       1,
							UserID:   1,
							Title:    "タイ旅行",
							IsPublic: true,
						},
						{
							ID:       2,
							UserID:   2,
							Title:    "ベトナム旅行",
							IsPublic: true,
						},
					}, nil
				},
			},
			wantTravels: []*Travel{
				{
					ID:       1,
					UserID:   1,
					Title:    "タイ旅行",
					IsPublic: true,
				},
				{
					ID:       2,
					UserID:   2,
					Title:    "ベトナム旅行",
					IsPublic: true,
				},
			},
		},
		{
			name: "公開Travelが0件",
			repository: &MockRepository{
				FindPublicTravelsFunc: func() ([]*Travel, error) {
					return []*Travel{}, nil
				},
			},
			wantTravels: []*Travel{},
		},
		{
			name: "Repositoryエラー",
			repository: &MockRepository{
				FindPublicTravelsFunc: func() ([]*Travel, error) {
					return nil, errors.New("repository error")
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &UseCase{
				repository: tt.repository,
			}

			got, err := uc.GetPublicTravels()

			if (err != nil) != tt.wantErr {
				t.Fatalf("GetPublicTravels() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.wantTravels) {
				t.Errorf("GetPublicTravels() = %v, want %v", got, tt.wantTravels)
			}
		})
	}
}
