package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Controller interface {
	GetUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
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

// GetUser godoc
// @Summary      Retrieve a user
// @Description  Fetches the details of a user by their ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  User
// @Failure      400  {object}  map[string]interface{}{"error": string}
// @Failure      404  {object}  map[string]interface{}{"error": string}
// @Router       /users/{id} [get]
func (c *ControllerImpl) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	// Validate the UUID format
	parse, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user UUID"})
		return
	}

	u, err := c.s.GetUserByID(parse)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, u)
}

// GetAllUsers godoc
// @Summary      Retrieve all users
// @Description  Fetches the details of all users in the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  []User
// @Failure		 500  {object}  map[string]interface{}{"error": string}
// @Router       /users [get]
func (c *ControllerImpl) GetAllUsers(ctx *gin.Context) {
	users, err := c.s.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Adds a new user to the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      User  true  "User input data"
// @Success      201   {object}  User
// @Failure      400   {object}  map[string]interface{}{"error": string}
// @Failure      500   {object}  map[string]interface{}{"error": string}
// @Router       /users [post]
func (c *ControllerImpl) CreateUser(ctx *gin.Context) {
	var userInput User

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := c.s.CreateUser(&userInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, u)
}

// UpdateUser godoc
// @Summary      Update user details
// @Description  Modifies details of an existing user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int   true  "User ID"
// @Param        user  body      User  true  "User updated data"
// @Success      200   {object}  User
// @Failure      400   {object}  map[string]interface{}{"error": string}
// @Failure      500   {object}  map[string]interface{}{"error": string}
// @Router       /users/{id} [put]
func (c *ControllerImpl) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	// Validate the UUID format
	_, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var userInput User

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInput.ID = id

	updatedUser, err := c.s.UpdateUser(&userInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Removes a user from the system by their ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      204  {object}  map[string]interface{}{"message": string}
// @Failure      400  {object}  map[string]interface{}{"error": string}
// @Failure      500  {object}  map[string]interface{}{"error": string}
// @Router       /users/{id} [delete]
func (c *ControllerImpl) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	// Validate the UUID format
	parse, err := uuid.Parse(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = c.s.DeleteUser(parse)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "User deleted"})
}
