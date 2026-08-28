package user

import (
	"errors"
	"net/http"

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

type RegisterRequest struct {
	Nickname             string  `json:"nickname" binding:"required"`
	Email                *string `json:"email"`
	Password             string  `json:"password" binding:"required"`
	PasswordConfirmation string  `json:"password_confirmation" binding:"required"`
}

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	input := RegisterInput{
		Nickname:             req.Nickname,
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}

	u, err := h.useCase.Register(input)
	if err != nil {
		if errors.Is(err, ErrEmailAlreadyUsed) {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}

		if errors.Is(err, ErrPasswordMismatch) {
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

	c.JSON(http.StatusCreated, gin.H{
		"id":       u.ID,
		"nickname": u.Nickname,
		"email":    u.Email,
	})
}
