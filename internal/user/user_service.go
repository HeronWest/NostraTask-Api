package user

type Service interface {
	GetUserByID(id uint) (*User, error)
	CreateUser(u *User) (*User, error)
	UpdateUser(u *User) (*User, error)
	DeleteUser(id uint) error
}

type ServiceImpl struct {
	r Repository
}

func NewUserService(r Repository) Service {
	return &ServiceImpl{r: r}
}

func (s *ServiceImpl) GetUserByID(id uint) (*User, error) {
	return s.r.FindByID(id)
}

func (s *ServiceImpl) CreateUser(u *User) (*User, error) {
	err := s.r.Create(u)
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

func (s *ServiceImpl) DeleteUser(id uint) error {
	return s.r.Delete(id)
}
