package travel

import "time"

const dateLayout = "2006-01-02"

type TravelResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date"`
	IsPublic  bool      `json:"is_public"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func toTravelResponse(t *Travel) TravelResponse {
	return TravelResponse{
		ID:        t.ID,
		UserID:    t.UserID,
		Title:     t.Title,
		StartDate: t.StartDate.Format(dateLayout),
		EndDate:   t.EndDate.Format(dateLayout),
		IsPublic:  t.IsPublic,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

type PublicTravelResponse struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	AuthorNickname string    `json:"author_nickname"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func toPublicTravelResponse(t *PublicTravel) PublicTravelResponse {
	return PublicTravelResponse{
		ID:             t.ID,
		Title:          t.Title,
		StartDate:      t.StartDate.Format(dateLayout),
		EndDate:        t.EndDate.Format(dateLayout),
		AuthorNickname: t.AuthorNickname,
		CreatedAt:      t.CreatedAt,
		UpdatedAt:      t.UpdatedAt,
	}
}
