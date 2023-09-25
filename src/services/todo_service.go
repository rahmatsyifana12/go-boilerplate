package services

import (
	"go-boilerplate/src/repositories"

	"github.com/sarulabs/di"
)

type TodoService interface {
	
}

type TodoServiceImpl struct {
	repository	*repositories.Repository
}

func NewTodoService(ioc di.Container) *TodoServiceImpl {
	return &TodoServiceImpl{
		repository: repositories.NewRepository(ioc),
	}
}
