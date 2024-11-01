package mock_repositories

import (
	"go-boilerplate/src/mocks"
	"go-boilerplate/src/repositories"

	"github.com/golang/mock/gomock"
)

var (
	UserRepository    *mocks.MockUserRepository
)

func NewMockRepository(ctrl *gomock.Controller) *repositories.Repository {
	UserRepository = mocks.NewMockUserRepository(ctrl)

	return &repositories.Repository{
		User: UserRepository,
	}
}