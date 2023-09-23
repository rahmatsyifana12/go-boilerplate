package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

	module := Module{}
	module.New(e)
    
    port, found := os.LookupEnv("PORT")
	if !found {
		port = "5000"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}