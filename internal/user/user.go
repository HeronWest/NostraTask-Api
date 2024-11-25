package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
