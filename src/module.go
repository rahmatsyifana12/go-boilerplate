package main

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/controllers"
	"go-boilerplate/src/databases"
	"go-boilerplate/src/repositories"
	"go-boilerplate/src/services"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type Module struct {}

func (m *Module) New(app *echo.Echo) {
	m.NewIOC()
	r := NewRoute(app)
	r.Init()
}

func (m *Module) NewIOC() di.Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add(
		di.Def{
			Name: constants.MYSQL,
			Build: func(ctn di.Container) (interface{}, error) {
				db, err := databases.NewMysqlClient()
				return db, err
			},
		},
		di.Def{
			Name: constants.CONTROLLER,
			Build: func(ctn di.Container) (interface{}, error) {
				return controllers.NewController(builder.Build()), nil
			},
		},
		di.Def{
			Name: constants.SERVICE,
			Build: func(ctn di.Container) (interface{}, error) {
				return services.NewService(builder.Build()), nil
			},
		},
		di.Def{
			Name: constants.REPOSITORY,
			Build: func(ctn di.Container) (interface{}, error) {
				return repositories.NewRepository(builder.Build()), nil
			},
		},
	)
	return builder.Build()
}