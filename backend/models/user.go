package models

import (
	"github.com/high-ping-devs/simple-chat-room/backend/auth"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null;size:255" validate:"nonzero,min=3,max=255"`
	Email    string `json:"email" gorm:"unique;not null;size:255" validate:"nonzero,email,min=3,max=255"`
	Password string `json:"password" gorm:"not null;size:255" validate:"nonzero,min=12,max=255"`
}

func (u *User) BeforeCreate(*gorm.DB) (err error) {
	u.Password = auth.HashAndSaltPassword(u.Password)
	return nil
}

func (u *User) PasswordMatch(password string) bool {
	return auth.PasswordMatch(u.Password, password)
}

func (u *User) Validate() error {
	return validator.Validate(u)
}
