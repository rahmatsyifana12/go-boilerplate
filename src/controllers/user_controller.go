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
			WithMessage("failed to bind params")
	}

	err = t.service.User.CreateUser(c, params)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusCreated).
		Send(c)
}

func (t *UserControllerImpl) GetUserByID(c echo.Context) (err error) {
	var (
		params	*dtos.GetUserByIDParams = new(dtos.GetUserByIDParams)
	)

	if err := (&echo.DefaultBinder{}).BindPathParams(c, params); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	data, err := t.service.User.GetUserByID(c, params.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}