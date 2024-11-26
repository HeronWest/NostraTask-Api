package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service interface {
	Login(email, password string) (string, error)
}

type ServiceImpl struct {
	r Repository
}

var jwtSecret = []byte("your_jwt_secret_key")

func NewAuthService(r Repository) Service {
	return &ServiceImpl{r: r}
}

func (s *ServiceImpl) Login(email, password string) (string, error) {
	auth, err := s.r.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"sub": auth.ID.String(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
