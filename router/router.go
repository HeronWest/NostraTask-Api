package router

import (
	"github.com/HeronWest/nostrataskapi/config"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(di *config.DependencyInjector) {
	logger := config.GetLogger("router")
	router := gin.Default()

	v1 := router.Group("/api/v1")

	di.Provide(func() *gin.RouterGroup {
		return v1
	})

	initializeRoutes(di)

	if err := router.Run(":8080"); err != nil {
		logger.Errorf("Erro ao iniciar o servidor: %v", err)
	} else {
		logger.Info("Servidor iniciado com sucesso.")
	}
}
