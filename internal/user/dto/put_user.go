package dto

type UpdateUserDTO struct {
	Name  string `json:"name" binding:"required,min=3,max=255"`
	Email string `json:"email" binding:"required,email" gorm:"unique"`
	Role  Role   `json:"role" binding:"omitempty,oneof=user admin"`
}
