package controllers

import (
	"fmt"
	"go-boilerplate/src/constants"
	"go-boilerplate/src/dtos"
	"go-boilerplate/src/pkg/responses"
	"go-boilerplate/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type UserController interface {
	CreateUser(c echo.Context) error
	GetUserByID(c echo.Context) error
}

type UserControllerImpl struct {
	service *services.Service
}

func NewUserController(ioc di.Container) *UserControllerImpl {
	return &UserControllerImpl{
        service: ioc.Get(constants.SERVICE).(*services.Service),
    }
}

func (t *UserControllerImpl) CreateUser(c echo.Context) (err error) {
	var (
		params	dtos.CreateUserRequest
	)

    if err = c.Bind(&params); err != nil {
		return responses.NewError().
			WithError(err).
			WithCode(http.StatusBadRequest).
			WithMessage("failed to bind parameters")
	}

	err = t.service.User.CreateUser(c, params)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusCreated).
		WithMessage("successfully created a new user").
		Send(c)
}

func (t *UserControllerImpl) GetUserByID(c echo.Context) (err error) {
	var (
		params	dtos.GetUserByIDParams
	)

	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters")
		return
	}

	fmt.Println(params)

	data, err := t.service.User.GetUserByID(c, params.UserID)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusOK).
		WithMessage("successfully retrieved a user").
		WithData(data).
		Send(c)
}