package main

import (
	"github.com/HeronWest/nostrataskapi/config"
	"github.com/HeronWest/nostrataskapi/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load .env file
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	db := config.InitializeDatabase()
	//redisClient := config.InitializeRedis()

	log.Println("Banco de dados conectado com sucesso:", db)

	//Inicialize Router
	router.InitializeRouter()
}
