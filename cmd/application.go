package cmd

import (
	"demo/infra/config"
	"demo/infra/connection"
	"fmt"

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
	port := config.App().Port
	e.Logger.Fatal(e.Start(":" + port))
}
