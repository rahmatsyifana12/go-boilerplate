package controllers

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/services"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type AuthController interface {
	Login(c echo.Context) error
}

type AuthControllerImpl struct {
	service	*services.Service
}

func NewAuthController(ioc di.Container) *AuthControllerImpl {
	return &AuthControllerImpl{
		service: ioc.Get(constants.SERVICE).(*services.Service),
	}
}

func (a *AuthControllerImpl) Login(c echo.Context) (err error) {

	return
}