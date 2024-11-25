package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

// Função para inicializar a conexão com o Redis
func InitializeRedis() *redis.Client {
	// Carregar as variáveis de ambiente
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	// Criar a string de conexão com o Redis
	addr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	// Conectar ao Redis
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	// Testar a conexão com o Redis
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Erro ao conectar ao Redis:", err)
	}

	return client
}
