package main

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/controllers"
	"go-boilerplate/src/middlewares"

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
	r.Auth()
	r.User()
}

func (r *Route) Test() {
	r.app.GET("/test", func(c echo.Context) error {
		return c.String(200, "Your application is working just fine :)")
	})
}

func (r *Route) User() {
	user := r.router.Group("users")
	user.POST("/", r.controller.User.CreateUser)
	user.GET("/:user_id", r.controller.User.GetUserByID, middlewares.AuthMiddleware)
	user.PATCH("/:user_id", r.controller.User.UpdateUser, middlewares.AuthMiddleware)
	user.DELETE("/:user_id", r.controller.User.DeleteUser, middlewares.AuthMiddleware)
}

func (r *Route) Auth() {
	auth := r.router.Group("auth")
	auth.POST("/login", r.controller.Auth.Login)
	auth.POST("/logout", r.controller.Auth.Logout, middlewares.AuthMiddleware)
}

func (r *Route) Todo() {
	todo := r.router.Group("todos")
	todo.POST("/", r.controller.Todo.CreateTodo, middlewares.AuthMiddleware)
}