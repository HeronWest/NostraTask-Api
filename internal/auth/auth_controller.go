package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller interface {
	Login(ctx *gin.Context)
}

type ControllerImpl struct {
	s Service
}

func NewAuthController(s Service) Controller {
	return &ControllerImpl{s: s}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c *ControllerImpl) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.s.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
