package auth

import (
	"errors"
	"github.com/HeronWest/nostrataskapi/internal/user"
	"gorm.io/gorm"
)

type Repository interface {
	FindByEmail(email string) (*Auth, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) FindByEmail(email string) (*Auth, error) {
	var u user.User
	err := r.db.Where("email = ?", email).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Converte os dados do usuário para a struct Auth
	auth := &Auth{
		ID:       u.ID, // Assumindo que o campo ID em User é do tipo uuid.UUID
		Email:    u.Email,
		Password: u.Password,
	}

	return auth, nil
}
