package services

import (
	"go-boilerplate/src/dtos"
	"go-boilerplate/src/models"
	"go-boilerplate/src/pkg/responses"
	"go-boilerplate/src/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type TodoService interface {
	CreateTodo(c echo.Context, claims dtos.AuthClaims, params dtos.CreateTodoRequest) error
	GetTodoByID(c echo.Context, claims dtos.AuthClaims, params dtos.TodoIDParams) (dtos.GetTodoByIDResponse, error)
	GetTodos(c echo.Context, claims dtos.AuthClaims) (dtos.GetTodosResponse, error)
}

type TodoServiceImpl struct {
	repository	*repositories.Repository
}

func NewTodoService(ioc di.Container) *TodoServiceImpl {
	return &TodoServiceImpl{
		repository: repositories.NewRepository(ioc),
	}
}

func (s *TodoServiceImpl) CreateTodo(c echo.Context, claims dtos.AuthClaims, params dtos.CreateTodoRequest) (err error) {
	user, err := s.repository.User.GetUserByID(c, claims.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = responses.NewError().
				WithError(err).
				WithCode(http.StatusBadRequest).
				WithMessage("Cannot find user with the given id")
			return
		}
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Cannot find user with the given id")
		return
	}

	newTodo := models.Todo{
		Title:   params.Title,
		Content: params.Content,
		UserID:  user.ID,
	}

	err = s.repository.Todo.CreateTodo(c, newTodo)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while creating todo into database")
		return
	}

	return
}

func (s *TodoServiceImpl) GetTodoByID(c echo.Context, claims dtos.AuthClaims, params dtos.TodoIDParams) (data dtos.GetTodoByIDResponse, err error) {
	todo, err := s.repository.Todo.GetTodoByID(c, params.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = responses.NewError().
				WithError(err).
				WithCode(http.StatusBadRequest).
				WithMessage("Cannot find todo with the given id")
			return
		}
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Cannot find todo with the given id")
		return
	}

	if todo.UserID != claims.UserID {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusUnauthorized).
			WithMessage("You are not authorized to view this todo")
		return
	}

	data.Todo = todo
	return
}

func (s *TodoServiceImpl) GetTodos(c echo.Context, claims dtos.AuthClaims) (data dtos.GetTodosResponse, err error) {
	todos, err := s.repository.Todo.GetTodosByUserID(c, claims.UserID)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while retrieving todos from database")
		return
	}

	data.Todos = todos
	return
}