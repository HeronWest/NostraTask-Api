package task

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// Controller interface defines the methods for handling tasks
type Controller interface {
	GetTaskById(c *gin.Context)
	GetAllTasksByUserId(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	DeleteUserTask(c *gin.Context)
	AddUserTask(c *gin.Context)
	GetAllUsersByTaskId(c *gin.Context)
	GetAllTaskHistoryByTaskId(c *gin.Context)
}

type ControllerImpl struct {
	s Service
}

func NewTaskController(s Service) Controller {
	return &ControllerImpl{s: s}
}

// GetTaskById godoc
// @Summary      Retrieve a task by ID
// @Description  Fetches the details of a task by its UUID, based on the user's claim
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Task ID (UUID)"
// @Success      200  {object}  Task
// @Failure      400  {object}  TaskResponse  "Invalid UUID format"
// @Failure      404  {object}  TaskResponse  "Task not found"
// @Router       /tasks/{id} [get]
func (c *ControllerImpl) GetTaskById(ctx *gin.Context) {
	id := ctx.Param("task_id")
	userID := ctx.MustGet("user_id").(uuid.UUID) // Extract user ID from the claim

	// Validate the UUID format
	parse, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, `{"error": "Invalid task UUID"}`)
		return
	}

	task, err := c.s.GetByID(parse, userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, `{"error": "`+err.Error()+`"}`)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// GetAllTasksByUserId godoc
// @Summary      Retrieve all tasks assigned to a user
// @Description  Fetches the details of all tasks assigned to the user from their claim
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Success      200  {object}  []Task
// @Failure      500  {object}  TaskResponse  "Internal server error"
// @Router       /tasks [get]
func (c *ControllerImpl) GetAllTasksByUserId(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uuid.UUID) // Extract user ID from the claim

	tasks, err := c.s.GetAllTasksByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, `{"error": "Failed to retrieve tasks"}`)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

// CreateTask godoc
// @Summary      Create a new task
// @Description  Adds a new task to the system, assigned to the user from their claim
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task  body      Task  true  "Task input data"
// @Success      201   {object}  Task
// @Failure      400   {object}  TaskResponse  "Invalid input data"
// @Failure      500   {object}  TaskResponse  "Internal server error"
// @Router       /tasks [post]
func (c *ControllerImpl) CreateTask(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uuid.UUID) // Extract user ID from the claim

	var task Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, `{"error": "`+err.Error()+`"}`)
		return
	}

	if err := c.s.Create(&task, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, `{"error": "Failed to create task, please try again"}`)
		return
	}

	ctx.JSON(http.StatusCreated, task)
}

// UpdateTask godoc
// @Summary      Update a task by ID
// @Description  Modifies an existing task's details, based on the user's claim
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Task ID (UUID)"
// @Param        task  body      Task    true  "Updated task data"
// @Success      200   {object}  Task
// @Failure      400   {object}  TaskResponse  "Invalid input or UUID format"
// @Failure      404   {object}  TaskResponse  "Task not found"
// @Failure      500   {object}  TaskResponse  "Internal server error"
// @Router       /tasks/{id} [put]
func (c *ControllerImpl) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("task_id")
	userID := ctx.MustGet("user_id").(uuid.UUID) // Extract user ID from the claim

	// Validate the UUID format
	parse, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, `{"error": "Invalid task ID"}`)
		return
	}

	var task Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	task.ID = parse

	if err := c.s.Update(&task, userID); err != nil {
		ctx.JSON(http.StatusNotFound, `{"error": "Task not found"}`)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// DeleteTask godoc
// @Summary      Delete a task by ID
// @Description  Removes a task from the system, based on the user's claim
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Task ID (UUID)"
// @Success      204  {object}  TaskResponse  "Task successfully deleted"
// @Failure      400  {object}  TaskResponse  "Invalid UUID format"
// @Failure      404  {object}  TaskResponse  "Task not found"
// @Failure      500  {object}  TaskResponse  "Internal server error"
// @Router       /tasks/{id} [delete]
func (c *ControllerImpl) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("task_id")
	userID := ctx.MustGet("user_id").(uuid.UUID) // Extract user ID from the claim

	// Validate the UUID format
	parse, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, `{"error": "Invalid task UUID"}`)
		return
	}

	if err := c.s.Delete(parse, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, `{"error": "Failed to delete task"}`)
		return
	}

	ctx.JSON(http.StatusNoContent, `{"message": "Task successfully deleted"}`)
}

// DeleteUserTask godoc
// @Summary      Remove a user from a task
// @Description  Removes a user from the task based on the provided task ID and user ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task_id   path      string  true  "Task ID (UUID)"
// @Param        user_id   path      string  true  "User ID (UUID)"
// @Success      200       {object}  TaskResponse  "User successfully removed from task"
// @Failure      400       {object}  TaskResponse  "Invalid UUID format"
// @Failure      404       {object}  TaskResponse  "Task or user not found"
// @Failure      500       {object}  TaskResponse  "Internal server error"
// @Router       /tasks/{task_id}/user/{user_id} [delete]
func (c *ControllerImpl) DeleteUserTask(ctx *gin.Context) {
	taskID := ctx.Param("task_id")
	userID := ctx.MustGet("user_id").(uuid.UUID) // Extract user ID from the claim
	removedUserID := ctx.Param("user_id")

	// Validate the UUID format for task and user
	taskUUID, err := uuid.Parse(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, `{"error": "Invalid task UUID"}`)
		return
	}

	userUUID, err := uuid.Parse(removedUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, `{"error": "Invalid user UUID"}`)
		return
	}

	if err := c.s.DeleteUserTask(taskUUID, userID, userUUID); err != nil {
		ctx.JSON(http.StatusInternalServerError, `{"error": "Failed to remove user from task"}`)
		return
	}

	ctx.JSON(http.StatusOK, `{"message": "User successfully removed from task"}`)
}

// AddUserTask godoc
// @Summary      Add a user to a task
// @Description  Adds a user to a task based on the provided task ID and user ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task_id   path      string  true  "Task ID (UUID)"
// @Param        user_id   path      string  true  "User ID (UUID)"
// @Success      200       {object}  TaskResponse  "User successfully added to task"
// @Failure      400       {object}  TaskResponse  "Invalid UUID format"
// @Failure      404       {object}  TaskResponse  "Task or user not found"
// @Failure      500       {object}  TaskResponse  "Internal server error"
// @Router       /tasks/{task_id}/user/{user_id} [post]
func (c *ControllerImpl) AddUserTask(ctx *gin.Context) {
	taskID := ctx.Param("task_id")
	userID := ctx.MustGet("user_id").(uuid.UUID) // Extract user ID from the claim
	addedUserID := ctx.Param("user_id")

	// Validate the UUID format for task and user
	taskUUID, err := uuid.Parse(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, `{"error": "Invalid task UUID"}`)
		return
	}

	userUUID, err := uuid.Parse(addedUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, `{"error": "Invalid user UUID"}`)
		return
	}

	if err := c.s.AddUserTask(taskUUID, userID, userUUID); err != nil {
		ctx.JSON(http.StatusInternalServerError, `{"error": "Failed to add user to task"}`)
		return
	}

	ctx.JSON(http.StatusOK, `{"message": "User successfully added to task"}`)
}

// GetAllUsersByTaskId godoc
// @Summary      Retrieve all users assigned to a task
// @Description  Fetches all users assigned to a specific task based on its UUID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task_id   path      string  true  "Task ID (UUID)"
// @Success      200       {object}  []user.User
// @Failure      400       {object}  TaskResponse  "Invalid UUID format"
// @Failure      500       {object}  TaskResponse  "Internal server error"
// @Router       /tasks/{task_id}/users [get]
func (c *ControllerImpl) GetAllUsersByTaskId(ctx *gin.Context) {
	taskID := ctx.Param("task_id")

	// Validate the UUID format for task
	taskUUID, err := uuid.Parse(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, `{"error": "Invalid task UUID"}`)
		return
	}

	users, err := c.s.GetAllUsersByTaskID(taskUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, `{"error": "Failed to retrieve users"}`)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// GetAllTaskHistoryByTaskId godoc
// @Summary      Retrieve task history by task ID
// @Description  Fetches the history of a specific task, including changes or updates over time
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task_id   path      string  true  "Task ID (UUID)"
// @Success      200       {object}  []TaskHistory
// @Failure      400       {object}  TaskResponse  "Invalid UUID format"
// @Failure      500       {object}  TaskResponse  "Internal server error"
// @Router       /tasks/{task_id}/history [get]
func (c *ControllerImpl) GetAllTaskHistoryByTaskId(ctx *gin.Context) {
	taskID := ctx.Param("task_id")

	// Validate the UUID format for task
	taskUUID, err := uuid.Parse(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, `{"error": "Invalid task UUID"}`)
		return
	}

	history, err := c.s.GetAllTaskHistoryByTaskID(taskUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, `{"error": "Failed to retrieve task history"}`)
		return
	}

	ctx.JSON(http.StatusOK, history)
}
