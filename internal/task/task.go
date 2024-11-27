package task

import (
	"github.com/google/uuid"
	"time"

	"github.com/HeronWest/nostrataskapi/internal/base"
	"github.com/lib/pq"
)

type Status string

const (
	StatusNew       Status = "New"
	StatusPending   Status = "Pending"
	StatusExecuting Status = "Executing"
	StatusFinished  Status = "Finished"
	StatusCancelled Status = "Cancelled"
)

type Permission string

const (
	PermissionView         Permission = "view"
	PermissionEdit         Permission = "edit"
	PermissionManageStatus Permission = "manage_status"
)

type Task struct {
	base.Base
	Title       string        `json:"title" binding:"required,min=3,max=255"`
	Description string        `json:"description" binding:"required"`
	Status      Status        `json:"status" binding:"required,oneof=New Pending Executing Finished Cancelled"`
	DueDate     time.Time     `json:"due_date" binding:"required"`
	Users       []TaskUser    `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	History     []TaskHistory `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}

type TaskUser struct {
	base.Base
	TaskID      uuid.UUID      `json:"task_id" gorm:"not null;index"`
	UserID      uuid.UUID      `json:"user_id" gorm:"not null;index"`
	Permissions pq.StringArray `json:"permissions" gorm:"type:text[]"` // Array-like for permissions
}

type TaskHistory struct {
	base.Base
	TaskID    uuid.UUID `json:"task_id" gorm:"not null;index"`
	UserID    uuid.UUID `json:"user_id" gorm:"not null;index"`
	Field     string    `json:"field" binding:"required"`
	OldValue  string    `json:"old_value"`
	NewValue  string    `json:"new_value"`
	ChangedAt time.Time `json:"changed_at"`
}

// Define o método para adicionar histórico de alterações
func (t *Task) AddHistory(userID uuid.UUID, field, oldValue, newValue string) TaskHistory {
	return TaskHistory{
		TaskID:    t.ID,
		UserID:    userID,
		Field:     field,
		OldValue:  oldValue,
		NewValue:  newValue,
		ChangedAt: time.Now(),
	}
}

// Func HasPermission
func (tu *TaskUser) HasPermission(p Permission) bool {
	for _, permission := range tu.Permissions {
		if Permission(permission) == p {
			return true
		}
	}
	return false
}
