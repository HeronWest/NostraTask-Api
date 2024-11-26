package config

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	// Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Printf("Error when trying to load the .env file: %v", err)
	}

	// Inicializar o logger
	logger = NewLogger("$config: ")
	logger.Infof("Logger initialized")

	// Inicializar o banco de dados
	var err error
	db, err = initializeDatabase()
	if err != nil {
		logger.Errorf("Error when trying to initialize the database: %v", err)
		return err
	}

	logger.Infof("Conection with the database established")
	return nil
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}

func GetDB() *gorm.DB {
	if db == nil {
		logger.Errorf("Database not initialized")
	}
	return db
}
