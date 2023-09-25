package repositories

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/models"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type TodoRepository interface {
	CreateTodo(c echo.Context, todo models.Todo) error
}

type TodoRepositoryImpl struct {
	db	*gorm.DB
}

func NewTodoRepository(ioc di.Container) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{
		db: ioc.Get(constants.POSTGRES).(*gorm.DB),
	}
}

func (r *TodoRepositoryImpl) CreateTodo(c echo.Context, todo models.Todo) error {
	err := r.db.Create(&todo).WithContext(c.Request().Context()).Error
	return err
}