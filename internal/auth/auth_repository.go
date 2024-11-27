package auth

import (
	"errors"
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
	var auth Auth

	// Consulta diretamente na tabela sem depender da estrutura User
	err := r.db.Table("users").Where("email = ?", email).First(&auth).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &auth, nil
}
