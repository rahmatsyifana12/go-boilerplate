package tests

import (
	"go-boilerplate/src/dtos"
	"go-boilerplate/src/mock"
	mock_repositories "go-boilerplate/src/mock/repositories"
	"go-boilerplate/src/models"
	"go-boilerplate/src/services"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserServiceSuite struct {
	suite.Suite
	ctx				 echo.Context
	ctrl             *gomock.Controller
	userService      services.UserService
}

func (s *UserServiceSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	module := mock.ModuleMock(s.ctrl)
	s.userService = services.NewUserService(module)
}

func (s *UserServiceSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestAffiliateService(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}

func (s *UserServiceSuite) TestCreateUser() {
    s.Run("Success", func() {
        s.SetupTest()
        mock_repositories.UserRepository.EXPECT().GetUserByUsername(gomock.Any(), "rahmat").Return(nil,nil)
        mock_repositories.UserRepository.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)

        err := s.userService.CreateUser(s.ctx, dtos.CreateUserRequest{ Username: "rahmat", Password: "rahmat" })

        s.Equal(nil, err)
    })
}

func (s *UserServiceSuite) TestGetUserByID() {
	s.Run("Success With No Error", func() {
		s.SetupTest()
		mock_repositories.UserRepository.EXPECT().GetUserByID(gomock.Any(), uint(1)).Return(&models.User{ Model: gorm.Model{ ID: 1 } }, nil)

		data, err := s.userService.GetUserByID(s.ctx, dtos.AuthClaims{UserID: 1}, dtos.UserIDParams{ID: 1})

		s.Equal(nil, err)
		s.Equal(dtos.GetUserByIDResponse{ User: models.User{ Model: gorm.Model{ ID: 1 } } }, data)
	})
}