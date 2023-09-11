package repositories

import (
	"go-boilerplate/src/constants"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx echo.Context) (error)
}

type UserRepositoryImpl struct {
	db	*gorm.DB
}

func NewUserRepository(ioc di.Container) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: ioc.Get(constants.MYSQL).(*gorm.DB),
	}
}

func (r *UserRepositoryImpl) CreateUser(ctx echo.Context) (err error) {
	return
}