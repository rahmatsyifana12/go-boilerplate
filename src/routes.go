package main

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/controllers"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type Route struct {
	app		*echo.Echo
	router	*echo.Group
	controller	*controllers.Controller
}

func NewRoute(app *echo.Echo, ioc di.Container) *Route {
	return &Route{
		app: 	app,
		router: app.Group(""),
		controller: ioc.Get(constants.CONTROLLER).(*controllers.Controller),
	}
}

func (r *Route) Init() {
	r.Test()
	r.User()
}

func (r *Route) Test() {
	r.app.GET("/test", func(c echo.Context) error {
		return c.String(200, "hello world")
	})
}

func (r *Route) User() {
	user := r.router.Group("user")
	user.POST("", r.controller.User.CreateUser)
	user.GET("/:userID", r.controller.User.GetUserByID)
}