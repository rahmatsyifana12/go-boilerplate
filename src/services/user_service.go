package services

import (
	"go-boilerplate/src/dtos"
	"go-boilerplate/src/repositories"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type UserService interface {
	CreateUser(ctx echo.Context) (error)
	GetUserByID(ctx echo.Context, userID uint) (dtos.GetUserByIDResponse, error)
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

func (s *UserServiceImpl) GetUserByID(ctx echo.Context, userID uint) (data dtos.GetUserByIDResponse, err error) {
	data  = dtos.GetUserByIDResponse{}

	user, err := s.repository.User.GetUserByID(ctx, userID)
	if err != nil {
		return
	}

	data.User = user
	return
}