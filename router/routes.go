package router

import (
	"github.com/HeronWest/nostrataskapi/internal/task"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		task.SetupRoutes(v1)
	}
}
