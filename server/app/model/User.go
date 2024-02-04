package model

import "gorm.io/gorm"

// add swagger
// @description 用户
type User struct {
	gorm.Model
	Nickname string `gorm:"unique"`
	Password string
	Email    string `gorm:"unique"`
	Phone    string `gorm:"unique"`
	Avatar   string
}
