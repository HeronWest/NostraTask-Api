package user

import "github.com/gin-gonic/gin"

func SetupRoutes(group *gin.RouterGroup, controller Controller) {
	group.POST("/user", controller.CreateUser)
	group.GET("/user/:id", controller.GetUser)
	group.PUT("/user/:id", controller.UpdateUser)
	group.DELETE("/user/:id", controller.DeleteUser)
}
