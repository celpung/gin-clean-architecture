package entity

import (
	"gorm.io/gorm"
)

type User struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"not null" json:"name"`
	Email          string `gorm:"unique" json:"email"`
	Phone          string `gorm:"unique" json:"phone"`
	Password       string `json:"password"`
	Role           uint   `gorm:"not null;default:1" json:"role"`
	Active         bool   `gorm:"not null;default:1" json:"active"`
	EmailConfirmed bool   `gorm:"not null;default:1" json:"email_confirmed"`
	gorm.Model
}
