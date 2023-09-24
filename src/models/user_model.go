package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username    string  `json:"username"`
	Password    string  `json:"password"`
	AccessToken string  `json:"access_token"`
    Todos       []Todo
}