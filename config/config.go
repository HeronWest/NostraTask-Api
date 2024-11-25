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
		log.Printf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Inicializar o logger
	logger = NewLogger("$config: ")
	logger.Infof("Logger inicializado com sucesso")

	// Inicializar o banco de dados
	var err error
	db, err = initializeDatabase()
	if err != nil {
		logger.Errorf("Erro ao inicializar o banco de dados: %v", err)
		return err
	}

	logger.Infof("Conexão com o banco de dados estabelecida com sucesso")
	return nil
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}

func GetDB() *gorm.DB {
	if db == nil {
		logger.Errorf("Banco de dados não inicializado")
	}
	return db
}
