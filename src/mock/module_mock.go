package mock

import (
	"go-boilerplate/src/constants"
	mock_repositories "go-boilerplate/src/mock/repositories"

	"github.com/golang/mock/gomock"
	"github.com/sarulabs/di"
)

func ModuleMock(ctrl *gomock.Controller) di.Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add(
		di.Def{
			Name: constants.Repository,
			Build: func(ctn di.Container) (interface{}, error) {
				return mock_repositories.NewMockRepository(ctrl), nil
			},
		},
	)
	return builder.Build()
}
