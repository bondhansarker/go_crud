package cmd

import (
	"fmt"
	"os"

	container "demo/app"
	"demo/app/http/middlewares"
	"demo/infra/config"
	"demo/infra/connection"
	"github.com/labstack/echo/v4"
)

// Execute executes the root command
func Execute() {
	config.LoadConfig()
	connection.ConnectDb()
	fmt.Println("about to start the application")
	runServer()
}

func runServer() {
	e := echo.New()
	if err := middlewares.Attach(e); err != nil {
		os.Exit(1)
	}
	container.Init(e)
	port := config.App().Port
	e.Logger.Fatal(e.Start(":" + port))
}
