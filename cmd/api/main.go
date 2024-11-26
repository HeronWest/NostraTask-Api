package main

import (
	"github.com/HeronWest/nostrataskapi/config"
	"github.com/HeronWest/nostrataskapi/router"
	"github.com/joho/godotenv"
	"log"
)

// @title NostraTask API
// @version 1.0
// @description API for task management
// @host localhost:8080
// @BasePath /api/v1
func main() {
	di := config.NewDependencyInjector()
	ap := config.NewApplicationBindings(di)

	if err := ap.InitializeBindings(); err != nil {
		panic("Error when trying to initialize application bindings")
	}

	if err := godotenv.Load("../../../.env"); err != nil {
		log.Fatal("Error when trying to load .env file")
	}

	if err := config.Init(); err != nil {
		panic(err)
	}

	// Obter logger e banco de dados configurados
	logger := config.GetLogger("$main: ")

	di.Provide(config.GetDB)

	logger.Infof("System is running with success!")

	// Inicializar o roteador
	router.InitializeRouter(di)
}
