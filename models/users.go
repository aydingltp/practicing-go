package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Surname  string `gorm:"not null" json:"surname"`
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Role     bool   `gorm:"not null" json:"role"`
}
