package main

import (
	"github.com/labstack/echo/v4"
)

type Route struct {
	app		*echo.Echo
	router	*echo.Group
}

func NewRoute(app *echo.Echo) *Route {
	return &Route{
		app: 	app,
		router: app.Group(""),
	}
}

func (r *Route) Init() {
	
}