package travel_plan

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/HisakeyT/backpacker-platform/internal/travel"
	"github.com/HisakeyT/backpacker-platform/internal/user"
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

	travelPlan, err := h.useCase.CreateTravelPlan(
		userID,
		uint(travelID),
		input,
	)

	if err != nil {
		if errors.Is(err, travel.ErrTravelNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		if errors.Is(err, user.ErrUserNotAuthorized) {
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

	c.JSON(http.StatusCreated, travelPlan)
}

type UpdateTravelPlanRequest struct {
	Date      *string `json:"date"`
	Place     *string `json:"place"`
	Content   *string `json:"content"`
	SortOrder *int    `json:"sort_order"`
}

func (h *Handler) UpdateTravelPlan(c *gin.Context) {
	userID := c.GetUint("userID")

	travelPlanID64, err := strconv.ParseUint(c.Param("travel_plan_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid travel_plan_id",
		})
		return
	}

	travelID64, err := strconv.ParseUint(c.Param("travel_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid travel_id",
		})
		return
	}

	travelID := uint(travelID64)
	travelPlanID := uint(travelPlanID64)

	var updateReq UpdateTravelPlanRequest
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	var date *time.Time
	if updateReq.Date != nil {
		parsed, err := time.Parse("2006-01-02", *updateReq.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid date",
			})
			return
		}

		date = &parsed
	}

	updateInput := UpdateTravelPlanInput{
		Date:      date,
		Place:     updateReq.Place,
		Content:   updateReq.Content,
		SortOrder: updateReq.SortOrder,
	}

	travelPlan, err := h.useCase.UpdateTravelPlan(userID, travelID, travelPlanID, updateInput)
	if err != nil {
		if errors.Is(err, travel.ErrTravelNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		if errors.Is(err, user.ErrUserNotAuthorized) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}

		if errors.Is(err, ErrTravelPlanNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		if errors.Is(err, ErrTravelPlanNotBelongToTravel) {
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

	c.JSON(http.StatusOK, travelPlan)
}

func (h *Handler) GetTravelPlans(c *gin.Context) {
	userID := c.GetUint("userID")

	travelID, err := strconv.ParseUint(c.Param("travel_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid travel_id",
		})
		return
	}

	travelPlans, err := h.useCase.GetTravelPlans(userID, uint(travelID))
	if err != nil {
		if errors.Is(err, travel.ErrTravelNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		if errors.Is(err, user.ErrUserNotAuthorized) {
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

	c.JSON(http.StatusOK, travelPlans)
}
