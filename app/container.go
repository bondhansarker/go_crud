package app

import (
	"demo/app/http/controllers"
	"demo/app/repository"
	"demo/app/service"
	"demo/infra/connection"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	db := connection.Db()

	// register all repositories
	userRepository := repository.NewUsersRepository(db)

	//register all services
	userService := service.NewUsersService(userRepository)

	//register all controllers
	controllers.NewUsersController(e, userService)

}
