package auth

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(group *gin.RouterGroup, controller Controller) {
	group.POST("/auth/login", controller.Login)
}
