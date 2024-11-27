package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service interface {
	Login(email, password string) (string, error)
	GenerateToken(userID uuid.UUID, role string) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
}

type ServiceImpl struct {
	r Repository
}

var jwtSecret = []byte("your_jwt_secret_key") // Use um segredo seguro aqui

// Claims define os dados contidos no token JWT
type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

// NewAuthService cria uma nova instância do serviço de autenticação
func NewAuthService(r Repository) Service {
	return &ServiceImpl{r: r}
}

// Login valida as credenciais do usuário e retorna um token JWT
func (s *ServiceImpl) Login(email, password string) (string, error) {
	auth, err := s.r.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Gerar o token JWT após login
	token, err := s.GenerateToken(auth.ID, auth.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GenerateToken gera um novo token JWT com base no userID e role
func (s *ServiceImpl) GenerateToken(userID uuid.UUID, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token válido por 24 horas

	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken valida um token JWT e retorna as claims
func (s *ServiceImpl) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}
