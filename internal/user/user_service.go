package user

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetUserByID(id uuid.UUID) (*User, error)
	GetAllUsers() ([]User, error)
	CreateUser(u *User) (*User, error)
	UpdateUser(u *User) (*User, error)
	DeleteUser(id uuid.UUID) error
}

type ServiceImpl struct {
	r Repository
}

func NewUserService(r Repository) Service {
	return &ServiceImpl{r: r}
}

func (s *ServiceImpl) GetUserByID(id uuid.UUID) (*User, error) {
	return s.r.FindByID(id)
}

func (s *ServiceImpl) GetAllUsers() ([]User, error) {
	return s.r.FindAll()
}

func (s *ServiceImpl) CreateUser(u *User) (*User, error) {
	// Criptografar a senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err // Retorna erro caso a criptografia falhe
	}

	// Substituir a senha pelo hash criptografado
	u.Password = string(hashedPassword)

	// Salvar o usuário no repositório
	err = s.r.Create(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *ServiceImpl) UpdateUser(u *User) (*User, error) {
	err := s.r.Update(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *ServiceImpl) DeleteUser(id uuid.UUID) error {
	return s.r.Delete(id)
}
