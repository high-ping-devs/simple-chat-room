package models

import (
	"github.com/high-ping-devs/simple-chat-room/backend/auth"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null;size:255"`
	Email    string `json:"email" gorm:"unique;not null;size:255"`
	Password string `json:"password" gorm:"not null;size:255"`
}

func (u *User) BeforeCreate(*gorm.DB) {
	u.Password = auth.HashAndSaltPassword(u.Password)
}

func (u *User) PasswordMatch(password string) bool {
	return auth.PasswordMatch(u.Password, password)
}
