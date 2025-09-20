package apps

import (
	"go-boilerplate/internal/apps/rest/handlers"
	"go-boilerplate/internal/constants"
	"go-boilerplate/internal/pkg/databases"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/internal/usecases"

	"github.com/sarulabs/di"
)

func NewIOC() di.Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add(
		di.Def{
			Name: constants.Postgres,
			Build: func(ctn di.Container) (interface{}, error) {
				db, err := databases.NewPostgresClient()
				return db, err
			},
		},
		di.Def{
			Name: constants.Redis,
			Build: func(ctn di.Container) (interface{}, error) {
				rdb, err := databases.NewRedisClient()
				return rdb, err
			},
		},
		di.Def{
			Name: constants.Controller,
			Build: func(ctn di.Container) (interface{}, error) {
				return handlers.NewController(ctn), nil
			},
		},
		di.Def{
			Name: constants.Usecase,
			Build: func(ctn di.Container) (interface{}, error) {
				return usecases.NewUsecase(ctn), nil
			},
		},
		di.Def{
			Name: constants.Repository,
			Build: func(ctn di.Container) (interface{}, error) {
				return repositories.NewRepository(ctn), nil
			},
		},
	)
	return builder.Build()
}
