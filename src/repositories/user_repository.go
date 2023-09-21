package repositories

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/models"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(c echo.Context, user models.User) error
	GetUserByID(c echo.Context, userID uint) (models.User, error)
	GetUserByUsername(c echo.Context, username string) (models.User, error)
}

type UserRepositoryImpl struct {
	db	*gorm.DB
}

func NewUserRepository(ioc di.Container) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: ioc.Get(constants.POSTGRES).(*gorm.DB),
	}
}

func (r *UserRepositoryImpl) CreateUser(c echo.Context, user models.User) (err error) {
	err = r.db.Create(&user).WithContext(c.Request().Context()).Error
	return
}

func (r *UserRepositoryImpl) GetUserByID(c echo.Context, userID uint) (user models.User, err error) {
	err = r.db.First(&user).Where("user_id = ?", userID).WithContext(c.Request().Context()).Error
	return
}

func (r *UserRepositoryImpl) GetUserByUsername(c echo.Context, username string) (user models.User, err error) {
	// err = r.db.First(&user).Where("username = ?", username).WithContext(c.Request().Context()).Error
	err = r.db.Select("id", "username").Where("username = ?", username).First(&user).WithContext(c.Request().Context()).Error
	return
}