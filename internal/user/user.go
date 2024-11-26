package user

import (
	"github.com/HeronWest/nostrataskapi/internal/base"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	base.Base
	Name  string `json:"name" binding:"required,min=3,max=255"`
	Email string `json:"email" binding:"required,email" gorm:"unique"`
	Role  Role   `json:"role" binding:"omitempty,oneof=user admin"`
}
