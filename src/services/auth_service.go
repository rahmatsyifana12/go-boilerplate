package services

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/repositories"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type AuthService interface {
	Login(c echo.Context) (err error)
}

type AuthServiceImpl struct {
	repository	*repositories.Repository
}

func NewAuthService(ioc di.Container) *AuthServiceImpl {
	return &AuthServiceImpl{
		repository: ioc.Get(constants.REPOSITORY).(*repositories.Repository),
	}
}

func (s *AuthServiceImpl) Login(c echo.Context) (err error) {
	return
}