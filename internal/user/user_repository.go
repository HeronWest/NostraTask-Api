package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindByID(id uuid.UUID) (*User, error)
	FindAll() ([]User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uuid.UUID) error
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) FindByID(id uuid.UUID) (*User, error) {
	var u User
	err := r.db.First(&u, id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *RepositoryImpl) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *RepositoryImpl) Create(user *User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryImpl) Update(user *User) error {
	// Use o id para garantir o update correto
	err := r.db.Model(&User{}).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryImpl) Delete(id uuid.UUID) error {
	err := r.db.Delete(&User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
