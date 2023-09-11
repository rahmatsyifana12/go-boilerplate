package repositories

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/models"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx echo.Context) (error)
	GetUserByID(ctx echo.Context, userID uint) (models.User, error)
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

func (r *UserRepositoryImpl) GetUserByID(ctx echo.Context, userID uint) (user models.User, err error) {
	err = r.db.First(&user, userID).Error
	return
}