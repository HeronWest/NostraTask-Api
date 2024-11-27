package task

import (
	"github.com/google/uuid"
)

type Service interface {
	GetByID(id uuid.UUID, userID uuid.UUID) (*Task, error)
	GetAllTasksByUserID(userID uuid.UUID) ([]Task, error)
	Create(task *Task, userID uuid.UUID) error
	Update(task *Task, userID uuid.UUID) error
	Delete(id uuid.UUID, userID uuid.UUID) error
	DeleteUserTask(taskID, userID uuid.UUID, removedUserID uuid.UUID) error
	AddUserTask(taskID, userID uuid.UUID, addUserID uuid.UUID) error
	GetAllUsersByTaskID(taskID uuid.UUID) ([]TaskUser, error)
	GetAllTaskHistoryByTaskID(taskID uuid.UUID) ([]TaskHistory, error)
}

type ServiceImpl struct {
	r Repository
}

func NewTaskService(r Repository) Service {
	return &ServiceImpl{r: r}
}

func (s *ServiceImpl) GetByID(id uuid.UUID, userID uuid.UUID) (*Task, error) {
	return s.r.FindByID(id, userID)
}

func (s *ServiceImpl) GetAllTasksByUserID(userID uuid.UUID) ([]Task, error) {
	return s.r.FindAllTasksByUserID(userID)
}

func (s *ServiceImpl) Create(task *Task, userID uuid.UUID) error {
	return s.r.Create(task, userID)
}

func (s *ServiceImpl) Update(task *Task, userID uuid.UUID) error {
	return s.r.Update(task, userID)
}

func (s *ServiceImpl) Delete(id uuid.UUID, userID uuid.UUID) error {
	return s.r.Delete(id, userID)
}

func (s *ServiceImpl) DeleteUserTask(taskID, userID uuid.UUID, removedUserID uuid.UUID) error {
	return s.r.DeleteUserTask(taskID, userID, removedUserID)
}

func (s *ServiceImpl) AddUserTask(taskID, userID uuid.UUID, addUserID uuid.UUID) error {
	return s.r.AddUserTask(taskID, userID, addUserID)
}

func (s *ServiceImpl) GetAllUsersByTaskID(taskID uuid.UUID) ([]TaskUser, error) {
	return s.r.FindAllUsersByTaskID(taskID)
}

func (s *ServiceImpl) GetAllTaskHistoryByTaskID(taskID uuid.UUID) ([]TaskHistory, error) {
	return s.r.FindAllTaskHistoryByTaskID(taskID)
}
