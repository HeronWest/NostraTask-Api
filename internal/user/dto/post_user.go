package dto

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type PostUserDTO struct {
	Name     string `json:"name" binding:"required,min=3,max=255"`
	Password string `json:"password" binding:"required,min=6,max=255"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Role     Role   `json:"role" binding:"omitempty,oneof=user admin"`
}
