package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint
	Username string
	Fullname string
	Password string
	Email    string
}
