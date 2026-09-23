package travel

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase *UseCase
}

func NewHandler(useCase *UseCase) *Handler {
	return &Handler{useCase: useCase}
}

type CreateTravelRequest struct {
	Title     string `json:"title" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
	IsPublic  bool   `json:"is_public"`
}

func (h *Handler) CreateTravel(c *gin.Context) {
	var req CreateTravelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.MustGet("userID").(uint)

	startDate, err := parseDate(req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid start date format",
		})
		return
	}

	endDate, err := parseDate(req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid end date format",
		})
		return
	}

	input := CreateTravelInput{
		Title:     req.Title,
		StartDate: startDate,
		EndDate:   endDate,
		IsPublic:  req.IsPublic,
	}

	travel, err := h.useCase.CreateTravel(userID, input)
	if err != nil {
		if errors.Is(err, ErrInvalidDateRange) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, travel)
}

func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

func (h *Handler) GetTravels(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	travels, err := h.useCase.GetTravels(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, travels)
}
