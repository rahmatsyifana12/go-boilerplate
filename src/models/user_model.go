package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username    string  `json:"username"`
	Password    string  `json:"-"`
	FullName    string  `json:"full_name"`
	PhoneNumber string  `json:"phone_number"`
	AccessToken string  `json:"access_token"`
    Todos       []Todo
}