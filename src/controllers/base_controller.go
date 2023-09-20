package controllers

import "github.com/sarulabs/di"

type Controller struct {	
	User	UserController
	Auth	AuthController
}

func NewController(ioc di.Container) *Controller {
	return &Controller{
		User: NewUserController(ioc),
		Auth: NewAuthController(ioc),

	}
}