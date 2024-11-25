package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

func InitializeRouter() {
	// Creating a new Gin router
	router := gin.Default()

	initializeRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
