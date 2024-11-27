package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// Middleware para validar o token JWT
func Middleware(service Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtém o token JWT do cabeçalho Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		// O formato do cabeçalho é "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(401, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		// Valida o token
		claims, err := service.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": fmt.Sprintf("Invalid token: %v", err)})
			c.Abort()
			return
		}

		// Adiciona as claims ao contexto para serem usadas nas requisições subsequentes
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		// Continua com o próximo manipulador
		c.Next()
	}
}
