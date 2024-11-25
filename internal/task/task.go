package task

import (
	"github.com/HeronWest/nostrataskapi/internal/user"
	"gorm.io/gorm"
	"time"
)

// TaskStatus define o enum para o status da tarefa
type Status string

const (
	StatusPending    Status = "PENDING"
	StatusInProgress Status = "IN_PROGRESS"
	StatusCompleted  Status = "COMPLETED"
	StatusArchived   Status = "ARCHIVED"
)

// Task representa uma tarefa no sistema
type Task struct {
	gorm.Model
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	Status        Status      `json:"status" gorm:"type:varchar(20);default:'PENDING'"`
	CreatedAt     time.Time   `json:"created_at"`
	Deadline      time.Time   `json:"deadline"`
	OwnerID       uint        `json:"owner_id"`
	Owner         user.User   `json:"owner" gorm:"foreignKey:OwnerID"`
	VisibleToAll  bool        `json:"visible_to_all" gorm:"default:true"`
	AssignedUsers []user.User `json:"assigned_users" gorm:"many2many:task_users;"`
}
