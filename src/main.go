package main

import (
	"fmt"
	"os"
	"strings"
	echo_middlewares "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

	e.Use(echo_middlewares.CORSWithConfig(echo_middlewares.CORSConfig{
		AllowOrigins: strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
	}))

	module := Module{}
	module.New(e)
    
    port, found := os.LookupEnv("PORT")
	if !found {
		port = "5000"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}