package auth

import "github.com/google/uuid"

type Auth struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
}
