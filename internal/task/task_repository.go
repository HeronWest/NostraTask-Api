package task

import (
	"github.com/HeronWest/nostrataskapi/internal/task/errors"
	"github.com/HeronWest/nostrataskapi/internal/user"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"log"
)

type Repository interface {
	FindByID(id uuid.UUID, userID uuid.UUID) (*Task, error)
	FindAllTasksByUserID(userID uuid.UUID) ([]Task, error)
	Create(task *Task, userID uuid.UUID) error
	Update(task *Task, userID uuid.UUID) error
	Delete(id uuid.UUID, userID uuid.UUID) error
	DeleteUserTask(taskID, userID uuid.UUID, removedUserID uuid.UUID) error
	AddUserTask(taskID, userID uuid.UUID, addUserID uuid.UUID) error
	FindAllUsersByTaskID(taskID uuid.UUID) ([]TaskUser, error)
	FindAllTaskHistoryByTaskID(taskID uuid.UUID) ([]TaskHistory, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{db: db}
}

// Create a new task
func (r *RepositoryImpl) Create(task *Task, userID uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Find user in users table by ID
		u := user.User{}
		if err := tx.First(&u, userID).Error; err != nil {
			log.Println("Error when trying to find user: ", err)
			return err
		}

		log.Println("User found: ", u)

		// Add user to task
		task.Users = append(task.Users, TaskUser{
			UserID:      u.ID,
			Permissions: pq.StringArray{string(PermissionView), string(PermissionEdit), string(PermissionManageStatus)},
		})

		// Create task
		if err := tx.Create(task).Error; err != nil {
			log.Println("Error when trying to create task: ", err)
			return err
		}

		// Create Task History
		th := task.AddHistory(u.ID, "Status", "", "New")
		if err := tx.Create(&th).Error; err != nil {
			log.Println("Error when trying to create task history: ", err)
			return err
		}

		// Commit implícito ao retornar nil
		return nil
	})
}

// Update a task
func (r *RepositoryImpl) Update(task *Task, userID uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Find user in users table by ID
		tu := TaskUser{}
		if err := tx.Where("task_id = ? AND user_id = ?", task.ID, userID).First(&tu).Error; err != nil {
			log.Println("Error when trying to find user: ", err)
			return err

		}

		// Update task
		if err := tx.Save(task).Error; err != nil {
			log.Println("Error when trying to update task: ", err)
			return err
		}

		// Create Task History
		th := task.AddHistory(tu.UserID, "Status", "", string(task.Status))
		if err := tx.Create(&th).Error; err != nil {
			log.Println("Error when trying to create task history: ", err)
			return err
		}

		// Commit implícito ao retornar nil
		return nil
	})
}

// Delete a task
func (r *RepositoryImpl) Delete(id uuid.UUID, userID uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Find user in users table by ID
		tu := TaskUser{}
		if err := tx.Where("task_id = ? AND user_id = ?", id, userID).First(&tu).Error; err != nil {
			log.Println("Error when trying to find user: ", err)
			return err
		}

		// Find task in tasks table by ID
		t := Task{}
		if err := tx.Preload("Users").First(&t, id).Error; err != nil {
			log.Println("Error when trying to find task: ", err)
			return err
		}

		// Check if user is allowed to delete task
		if !tu.HasPermission(PermissionManageStatus) {
			return errors.PermissionDeniedError{Message: "User does not have permission to delete task"}
		}

		// Delete task
		if err := tx.Delete(&t).Error; err != nil {
			log.Println("Error when trying to delete task: ", err)
			return err
		}

		// Commit implícito ao retornar nil
		return nil
	})
}

// Find a task by ID
func (r *RepositoryImpl) FindByID(id uuid.UUID, userID uuid.UUID) (*Task, error) {
	var t Task
	// Buscar a tarefa e carregar os relacionamentos 'Users' e 'History'
	if err := r.db.Preload("Users").Preload("History").First(&t, id).Error; err != nil {
		log.Println("Error when trying to find task: ", err)
		return nil, err
	}

	// Verificar se o usuário está vinculado à tarefa
	userFound := false
	for _, taskUser := range t.Users {
		if taskUser.UserID == userID {
			userFound = true
			break
		}
	}

	// Se o usuário não estiver vinculado à tarefa, retornar o erro
	if userFound == false {
		return nil, errors.UserNotFoundError{Message: "User not found in task"}
	}

	// Se a tarefa for encontrada e o usuário estiver vinculado, retornar a tarefa
	return &t, nil
}

// Find all tasks by user ID
func (r *RepositoryImpl) FindAllTasksByUserID(userID uuid.UUID) ([]Task, error) {
	var tasks []Task
	if err := r.db.
		Preload("Users").                                     // Precarrega os usuários
		Preload("History").                                   // Precarrega o histórico
		Joins("JOIN task_users tu ON tu.task_id = tasks.id"). // Faz a junção com a tabela intermediária
		Where("tu.user_id = ?", userID).                      // Filtra pelo user_id na tabela intermediária
		Find(&tasks).Error; err != nil {
		log.Println("Error when trying to find tasks: ", err)
		return nil, err
	}

	return tasks, nil
}

// Delete a user from a task
func (r *RepositoryImpl) DeleteUserTask(taskID, userID, removedUserID uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Find user in users table by ID
		tu := TaskUser{}
		if err := tx.Where("task_id = ? AND user_id = ?", taskID, userID).First(&tu).Error; err != nil {
			log.Println("Error when trying to find user: ", err)
			return err
		}

		// Find task in tasks table by ID
		t := Task{}
		if err := tx.Preload("Users").First(&t, taskID).Error; err != nil {
			log.Println("Error when trying to find task: ", err)
			return err
		}

		// Check if user is allowed to delete user from task
		if !tu.HasPermission(PermissionManageStatus) {
			return errors.PermissionDeniedError{Message: "User does not have permission to delete user from task"}
		}

		// Delete user from task
		if err := tx.Where("task_id = ? AND user_id = ?", taskID, removedUserID).Delete(&TaskUser{}).Error; err != nil {
			log.Println("Error when trying to delete user from task: ", err)
			return err
		}

		// Update task
		if err := tx.Save(&t).Error; err != nil {
			log.Println("Error when trying to update task: ", err)
			return err
		}

		// Create Task History
		th := t.AddHistory(tu.UserID, "User", removedUserID.String(), "")
		if err := tx.Create(&th).Error; err != nil {
			log.Println("Error when trying to create task history: ", err)
			return err
		}

		// Commit implícito ao retornar nil
		return nil
	})
}

// Add a user to a task
func (r *RepositoryImpl) AddUserTask(taskID, userID uuid.UUID, addUserID uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Find user in users table by ID
		u := user.User{}
		if err := tx.First(&u, addUserID).Error; err != nil {
			log.Println("Error when trying to find user: ", err)
			return err
		}

		// Find user in UserTask
		tu := TaskUser{}
		if err := tx.Where("task_id = ? AND user_id = ?", taskID, userID).First(&tu).Error; err != nil {
			log.Println("Error when trying to find user: ", err)
			return err
		}

		// Find task in tasks table by ID
		t := Task{}
		if err := tx.Preload("Users").First(&t, taskID).Error; err != nil {
			log.Println("Error when trying to find task: ", err)
			return err
		}

		// Check if user is allowed to add user to task
		if !tu.HasPermission(PermissionManageStatus) {
			return errors.PermissionDeniedError{Message: "User does not have permission to add user to task"}
		}

		// Add user to task
		t.Users = append(t.Users, TaskUser{
			UserID:      u.ID,
			Permissions: pq.StringArray{string(PermissionView), string(PermissionEdit)},
		})

		// Update task
		if err := tx.Save(&t).Error; err != nil {
			log.Println("Error when trying to update task: ", err)
			return err
		}

		// Create Task History
		th := t.AddHistory(tu.UserID, "User", "", addUserID.String())
		if err := tx.Create(&th).Error; err != nil {
			log.Println("Error when trying to create task history", err)
			return err
		}
		return nil
	})
}

// Find all users by task ID
func (r *RepositoryImpl) FindAllUsersByTaskID(taskID uuid.UUID) ([]TaskUser, error) {
	var users []TaskUser
	if err := r.db.Where("task_id = ?", taskID).Find(&users).Error; err != nil {
		log.Println("Error when trying to find users: ", err)
		return nil, err
	}
	return users, nil
}

// Find all task history by task ID
func (r *RepositoryImpl) FindAllTaskHistoryByTaskID(taskID uuid.UUID) ([]TaskHistory, error) {
	var history []TaskHistory
	if err := r.db.Where("task_id = ?", taskID).Find(&history).Error; err != nil {
		log.Println("Error when trying to find history: ", err)
		return nil, err
	}
	return history, nil
}
