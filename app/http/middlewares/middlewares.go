package middlewares

import (
	"demo/infra/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Attach(e *echo.Echo) error {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())

	e.Use(JWTWithConfig(JWTConfig{
		Skipper: func(context echo.Context) bool {
			switch context.Request().URL.Path {
			case "/users/login",
				"/users/signup":
				return true
			default:
				return false
			}
		},
		SigningKey: []byte(config.Jwt().AccessTokenSecret),
		ContextKey: config.Jwt().ContextKey,
	}))
	return nil
}
