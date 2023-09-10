package main

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/databases"

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
	)
	return builder.Build()
}