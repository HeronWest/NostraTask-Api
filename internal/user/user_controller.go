package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type ControllerImpl struct {
	s Service
}

func NewUserController(s Service) Controller {
	return &ControllerImpl{s: s}
}

// GetUser - Obtém o usuário com base no ID
func (c *ControllerImpl) GetUser(ctx *gin.Context) {
	// Pega o ID do usuário da URL
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	u, err := c.s.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Retorna os dados do usuário
	ctx.JSON(http.StatusOK, u)
}

// CreateUser - Cria um novo usuário
func (c *ControllerImpl) CreateUser(ctx *gin.Context) {
	var userInput User

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	u, err := c.s.CreateUser(&userInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Retorna o usuário criado
	ctx.JSON(http.StatusCreated, u)
}

// UpdateUser - Atualiza os dados do usuário
func (c *ControllerImpl) UpdateUser(ctx *gin.Context) {
	_, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var userInput User

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedUser, err := c.s.UpdateUser(&userInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Retorna o usuário atualizado
	ctx.JSON(http.StatusOK, updatedUser)
}

// DeleteUser - Deleta um usuário
func (c *ControllerImpl) DeleteUser(ctx *gin.Context) {
	// Pega o ID do usuário da URL
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Chama o serviço para deletar o usuário
	err = c.s.DeleteUser(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	// Retorna resposta de sucesso
	ctx.JSON(http.StatusNoContent, gin.H{"message": "User deleted"})
}
