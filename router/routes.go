package router

import (
	"github.com/HeronWest/nostrataskapi/config"
	docs "github.com/HeronWest/nostrataskapi/docs"
	"github.com/HeronWest/nostrataskapi/internal/user"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(di *config.DependencyInjector) {
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	di.Invoke(user.SetupRoutes)
	di.Invoke(initializeSwagger)

	// Swagg docs

}

func initializeSwagger(router *gin.RouterGroup) {
	// Swagger

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
