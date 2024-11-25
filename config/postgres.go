package config

import (
	"fmt"
	"github.com/HeronWest/nostrataskapi/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// Função para inicializar a conexão com o banco de dados
func initializeDatabase() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	// Carregar as variáveis de ambiente
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Formatar a string de conexão com o PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

	// Conectar ao banco de dados
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Erro ao conectar ao banco de dados:", err)
		return nil, err
	}

	err = db.AutoMigrate(
		&user.User{},
	)
	if err != nil {
		logger.Error("Erro ao realizar migrações:", err)
		return nil, err
	}

	logger.Info("Conexão com o banco de dados estabelecida.")

	return db, nil
}
