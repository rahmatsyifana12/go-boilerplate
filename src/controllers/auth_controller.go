package controllers

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/dtos"
	"go-boilerplate/src/pkg/responses"
	"go-boilerplate/src/services"
	"net/http"

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
	var (
		params	dtos.LoginRequest
	)

	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters")
		return
	}

	res, err := a.service.Auth.Login(c, params)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusOK).
		WithMessage("Successfully logged in").
		WithData(res).
		Send(c)
}