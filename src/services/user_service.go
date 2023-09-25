package services

import (
	"go-boilerplate/src/dtos"
	"go-boilerplate/src/models"
	"go-boilerplate/src/pkg/responses"
	"go-boilerplate/src/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(c echo.Context, createUserRequest dtos.CreateUserRequest) error
	GetUserByID(c echo.Context, userID uint) (dtos.GetUserByIDResponse, error)
}

type UserServiceImpl struct {
	repository	*repositories.Repository
}

func NewUserService(ioc di.Container) *UserServiceImpl {
	return &UserServiceImpl{
		repository: repositories.NewRepository(ioc),
	}
}

func (s *UserServiceImpl) CreateUser(c echo.Context, createUserRequest dtos.CreateUserRequest) (err error) {
	user, err := s.repository.User.GetUserByUsername(c, createUserRequest.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage(err.Error())
		return
	}

	if user.ID != 0 {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusBadRequest).
			WithMessage("Account with the same username already exists")
		return
	}

	passBytes := []byte(createUserRequest.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to hash password")
		return
	}

	newUser := models.User{
		Username: createUserRequest.Username,
		Password: string(hashedPassword),
	}

	err = s.repository.User.CreateUser(c, newUser)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage(err.Error())
		return
	}

	return
}

func (s *UserServiceImpl) GetUserByID(c echo.Context, userID uint) (data dtos.GetUserByIDResponse, err error) {
	user, err := s.repository.User.GetUserByID(c, userID)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Cannot find user with the given id")
		return
	}

	data.User = user
	return
}