package main

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/controllers"
	"go-boilerplate/src/pkg/databases"
	"go-boilerplate/src/pkg/utils"
	"go-boilerplate/src/repositories"
	"go-boilerplate/src/services"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type Module struct{}

func (m *Module) New(app *echo.Echo) {
	ioc := m.NewIOC()

	r := NewRoute(app, ioc)
	r.Init()
}

func (m *Module) NewIOC() di.Container {
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
				return controllers.NewController(builder.Build()), nil
			},
		},
		di.Def{
			Name: constants.Service,
			Build: func(ctn di.Container) (interface{}, error) {
				return services.NewService(builder.Build()), nil
			},
		},
		di.Def{
			Name: constants.Repository,
			Build: func(ctn di.Container) (interface{}, error) {
				return repositories.NewRepository(builder.Build()), nil
			},
		},
		di.Def{
			Name: constants.Util,
			Build: func(ctn di.Container) (interface{}, error) {
				return utils.NewUtil(), nil
			},
		},
	)
	return builder.Build()
}
