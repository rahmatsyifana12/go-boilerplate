package controllers

import (
	"fmt"
	"go-boilerplate/src/constants"
	"go-boilerplate/src/services"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type UserController interface {
	CreateUser(ctx echo.Context) (error)
}

type UserControllerImpl struct {
	service *services.Service
}

func NewUserController(ioc di.Container) *UserControllerImpl {
	return &UserControllerImpl{
        service: ioc.Get(constants.SERVICE).(*services.Service),
    }
}

func (c *UserControllerImpl) CreateUser(ctx echo.Context) (err error) {
    fmt.Println("CreateUser")
    c.service.User.CreateUser(ctx)
    return
}