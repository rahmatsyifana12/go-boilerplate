package repositories

import (
	"go-boilerplate/src/constants"

	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type TodoRepository interface {
	
}

type TodoRepositoryImpl struct {
	db	*gorm.DB
}

func NewTodoRepository(ioc di.Container) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{
		db: ioc.Get(constants.POSTGRES).(*gorm.DB),
	}
}