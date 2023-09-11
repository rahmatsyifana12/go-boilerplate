package controllers

import (
	"fmt"
	"go-boilerplate/src/constants"
	"go-boilerplate/src/dtos"
	"go-boilerplate/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type UserController interface {
	CreateUser(ctx echo.Context) (error)
	GetUserByID(ctx echo.Context) (error)
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

func (c *UserControllerImpl) GetUserByID(ctx echo.Context) (err error) {
	var (
		params	*dtos.GetUserByIDParams = new(dtos.GetUserByIDParams)
	)

	if err = ctx.Bind(params); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	data, err := c.service.User.GetUserByID(ctx, params.UserID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, data)
}