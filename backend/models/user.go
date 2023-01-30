package models

import (
	"net/mail"

	"github.com/high-ping-devs/simple-chat-room/backend/auth"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null;size:255" validate:"nonzero,min=3,max=255"`
	Email    string `json:"email" gorm:"unique;not null;size:255" validate:"nonzero,min=3,max=255,email"`
	Password string `json:"password" gorm:"not null;size:255" validate:"nonzero,min=12,max=255"`
}

func EmailValidator(v interface{}, param string) error {
	_, err := mail.ParseAddress(v.(string))
	return err
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = auth.HashAndSaltPassword(u.Password)
	return
}

func (u *User) PasswordMatch(password string) bool {
	return auth.PasswordMatch(password, u.Password)
}

func (u *User) Validate() error {
	return validator.Validate(u)
}
