package services

import (
	"context"
	"go-boilerplate/src/constants"
	"go-boilerplate/src/dtos"
	"go-boilerplate/src/models"
	"go-boilerplate/src/pkg/responses"
	"go-boilerplate/src/repositories"
	"net/http"

	"github.com/sarulabs/di"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, params dtos.CreateUserRequest) (err error)
	GetUserByID(ctx context.Context, claims dtos.AuthClaims, params dtos.UserIDParams) (data dtos.GetUserByIDResponse, err error)
	UpdateUser(ctx context.Context, claims dtos.AuthClaims, params dtos.UpdateUserParams) (err error)
	DeleteUser(ctx context.Context, claims dtos.AuthClaims, params dtos.UserIDParams) (err error)
}

type UserServiceImpl struct {
	repository	*repositories.Repository
}

func NewUserService(ioc di.Container) *UserServiceImpl {
	return &UserServiceImpl{
		repository: ioc.Get(constants.Repository).(*repositories.Repository),
	}
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, params dtos.CreateUserRequest) (err error) {
	user, err := s.repository.User.GetUserByUsername(ctx, params.Username)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while retrieving user by username from database")
		return
	}

	if user != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusBadRequest).
			WithMessage("Account with the same username already exists")
		return
	}

	passBytes := []byte(params.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to hash password")
		return
	}

	newUser := models.User{
		Username: params.Username,
		Password: string(hashedPassword),
		FullName: params.FullName,
		PhoneNumber: params.PhoneNumber,
	}

	err = s.repository.User.CreateUser(ctx, newUser)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while creating user into database")
		return
	}

	return
}

func (s *UserServiceImpl) GetUserByID(ctx context.Context, claims dtos.AuthClaims, params dtos.UserIDParams) (data dtos.GetUserByIDResponse, err error) {
	user, err := s.repository.User.GetUserByID(ctx, params.ID)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while retrieving user by id from database")
		return
	}

	if user == nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusBadRequest).
			WithMessage("Cannot find user with the given id")
		return
	}

	if user.ID != claims.UserID {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusUnauthorized).
			WithMessage("You are not authorized to view this user")
		return
	}

	data.User = *user
	return
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, claims dtos.AuthClaims, params dtos.UpdateUserParams) (err error) {
	user, err := s.repository.User.GetUserByID(ctx, params.ID)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while retrieving user by id from database")
		return
	}

	if user == nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusBadRequest).
			WithMessage("Cannot find user with the given id")
		return
	}
	
	if user.ID != claims.UserID {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusUnauthorized).
			WithMessage("You are not authorized to update this user")
		return
	}

	user.FullName = params.FullName
	user.PhoneNumber = params.PhoneNumber

	err = s.repository.User.UpdateUser(ctx, *user)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Cannot update user")
		return
	}

	return
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, claims dtos.AuthClaims, params dtos.UserIDParams) (err error) {
	user, err := s.repository.User.GetUserByID(ctx, params.ID)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while retrieving user by id from database")
		return
	}

	if user == nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusBadRequest).
			WithMessage("Cannot find user with the given id")
		return
	}

	if user.ID != claims.UserID {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusUnauthorized).
			WithMessage("You are not authorized to delete this user")
		return
	}

	err = s.repository.User.DeleteUser(ctx, *user)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while deleting user from database")
		return
	}

	return
}