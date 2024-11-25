package user

import (
	"github.com/HeronWest/nostrataskapi/internal/base"
)

type User struct {
	base.Base
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Role  string `json:"role"`
}
