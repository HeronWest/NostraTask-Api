package user

import (
	"github.com/HeronWest/nostrataskapi/internal/auth"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(group *gin.RouterGroup, controller Controller, as auth.Service) {

	group.POST("/user", auth.Middleware(as), controller.CreateUser)
	group.GET("/user/:id", auth.Middleware(as), controller.GetUser)
	group.GET("/user", auth.Middleware(as), controller.GetAllUsers)
	group.PUT("/user/:id", auth.Middleware(as), controller.UpdateUser)
	group.DELETE("/user/:id", auth.Middleware(as), controller.DeleteUser)
}
