package task

import "github.com/gin-gonic/gin"

func SetupRoutes(group *gin.RouterGroup) {
	group.POST("/task", CreateTask)
	group.GET("/task", GetTask)
	group.PUT("/task/:id", UpdateTask)
	group.DELETE("/task/:id", DeleteTask)
}
