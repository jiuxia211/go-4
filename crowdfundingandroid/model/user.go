package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	Telephone      string
	Class          uint //0管理员  1一般用户 2提出众筹者
	Money          int64
}

func (user *User) SetPassword(password string) error {
	passwordDigest, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	} else {
		user.PasswordDigest = string(passwordDigest)
		return nil
	}
}
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
