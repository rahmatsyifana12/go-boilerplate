package services

import (
	"go-boilerplate/src/repositories"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type UserService interface {
	CreateUser(ctx echo.Context) (error)
}

type UserServiceImpl struct {
	repository	*repositories.Repository
}

func NewUserService(ioc di.Container) *UserServiceImpl {
	return &UserServiceImpl{
		repository: repositories.NewRepository(ioc),
	}
}

func (s *UserServiceImpl) CreateUser(ctx echo.Context) (err error) {
	return
}