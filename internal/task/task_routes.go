package task

import (
	"github.com/HeronWest/nostrataskapi/internal/auth"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(group *gin.RouterGroup, controller Controller, as auth.Service) {
	group.POST("/task", auth.Middleware(as), controller.CreateTask)

	group.GET("/task/:task_id", auth.Middleware(as), controller.GetTaskById)
	group.GET("/tasks", auth.Middleware(as), controller.GetAllTasksByUserId)

	group.PUT("/task/:task_id", auth.Middleware(as), controller.UpdateTask)
	group.DELETE("/task/:task_id", auth.Middleware(as), controller.DeleteTask)
	group.DELETE("/task/:task_id/user/:user_id", auth.Middleware(as), controller.DeleteUserTask)

	group.POST("/task/:task_id/user/:user_id", auth.Middleware(as), controller.AddUserTask)

	group.GET("/task/:task_id/users", auth.Middleware(as), controller.GetAllUsersByTaskId)
	group.GET("/task/:task_id/history", auth.Middleware(as), controller.GetAllTaskHistoryByTaskId)
}
