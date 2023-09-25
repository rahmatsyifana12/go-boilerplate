package services

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/dtos"
	"go-boilerplate/src/pkg/helpers"
	"go-boilerplate/src/pkg/responses"
	"go-boilerplate/src/repositories"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(c echo.Context, params dtos.LoginRequest) (dtos.LoginResponse, error)
	Logout(c echo.Context, authClaims dtos.AuthClaims) error
}

type AuthServiceImpl struct {
	repository	*repositories.Repository
}

func NewAuthService(ioc di.Container) *AuthServiceImpl {
	return &AuthServiceImpl{
		repository: ioc.Get(constants.REPOSITORY).(*repositories.Repository),
	}
}

func (s *AuthServiceImpl) Login(c echo.Context, params dtos.LoginRequest) (dtos.LoginResponse, error) {
	var (
		res	dtos.LoginResponse
		err	error
	)
	user, err := s.repository.User.GetUserByUsername(c, params.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, responses.NewError().
				WithError(err).
				WithCode(http.StatusBadRequest).
				WithMessage("Cannot find user with the given username")
		}
		return res, responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while retrieving user from database")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		return res, responses.NewError().
			WithError(err).
			WithCode(http.StatusBadRequest).
			WithMessage("Incorrect password")
	}

	tokenExpireDuration := (time.Hour * 24)
	currentTime := time.Now()

	token, err := helpers.GenerateJWTString(dtos.AuthClaims{
		UserID:       user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(tokenExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(currentTime),
		},
	})
	if err != nil {
		return res, responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while generating JWT")
	}

	user.AccessToken = token

	err = s.repository.User.UpdateUser(c, user)
	if err != nil {
		return res, responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while updating users access token into database")
	}

	res.AccessToken = token
	return res, nil
}

func (s *AuthServiceImpl) Logout(c echo.Context, authClaims dtos.AuthClaims) error {
	user, err := s.repository.User.GetUserByID(c, authClaims.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.NewError().
				WithError(err).
				WithCode(http.StatusBadRequest).
				WithMessage("Cannot find user with the given id")
		}
		return responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while retrieving user from database")
	}

	user.AccessToken = ""

	err = s.repository.User.UpdateUser(c, user)
	if err != nil {
		return responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while updating users access token into database")
	}

	return nil
}