package main

import (
	"github.com/HeronWest/nostrataskapi/config"
	"github.com/HeronWest/nostrataskapi/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	di := config.NewDependencyInjector()
	ap := config.NewApplicationBindings(di)

	if err := ap.InitializeBindings(); err != nil {
		panic("Erro ao inicializar as dependências da aplicação")
	}

	if err := godotenv.Load("../../../.env"); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	if err := config.Init(); err != nil {
		panic(err)
	}

	// Obter logger e banco de dados configurados
	logger := config.GetLogger("$main: ")

	di.Provide(config.GetDB)

	logger.Infof("Sistema inicializado com sucesso!")

	// Inicializar o roteador
	router.InitializeRouter(di)
}
