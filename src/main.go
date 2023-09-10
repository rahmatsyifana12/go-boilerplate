package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
    app := echo.New()

	module := Module{}
	module.New(app)
    
    port, found := os.LookupEnv("PORT")
	if !found {
		port = "5000"
	}

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", port)))
}