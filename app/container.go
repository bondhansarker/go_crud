package app

import (
	"context"

	"demo/app/http/controllers"
	"demo/app/repository"
	"demo/app/service"
	"demo/infra/connection"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	baseContext := context.Background()
	db := connection.Db()

	// register all repositories
	userRepository := repository.NewUsersRepository(db)

	//register all services
	userService := service.NewUsersService(userRepository)
	tokenService := service.NewTokenService(baseContext, userRepository)
	authService := service.NewAuthService(userRepository, tokenService)

	//register all controllers
	controllers.NewUsersController(e, userService)
	controllers.NewAuthController(e, authService, userService)
}
