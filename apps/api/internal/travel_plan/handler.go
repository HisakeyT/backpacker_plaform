package travel_plan

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase *UseCase
}

func NewHandler(useCase *UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type CreateTravelPlanRequest struct {
	Date      string `json:"date"`
	Place     string `json:"place"`
	Content   string `json:"content"`
	SortOrder int    `json:"sort_order"`
}

func (h *Handler) CreateTravelPlan(c *gin.Context) {
	userID := c.GetUint("userID")

	travelID, err := strconv.ParseUint(c.Param("travel_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid travel_id",
		})
		return
	}

	var req CreateTravelPlanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid date",
		})
		return
	}

	input := CreateTravelPlanInput{
		Date:      date,
		Place:     req.Place,
		Content:   req.Content,
		SortOrder: req.SortOrder,
	}

	if err := h.useCase.CreateTravelPlan(
		userID,
		uint(travelID),
		input,
	); err != nil {
		if errors.Is(err, ErrUserNotAuthorized) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
