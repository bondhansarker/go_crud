package controllers

import (
	"net/http"

	"demo/app/serializers"
	"demo/app/service"
	"demo/infra/errors"
	"github.com/labstack/echo/v4"
)

type auth struct {
	authService service.IAuth
	userService service.IUsers
}

// NewAuthController will initialize the controllers
func NewAuthController(e *echo.Echo, authService service.IAuth, userService service.IUsers) {
	ac := &auth{
		authService: authService,
		userService: userService,
	}

	e.POST("/users/login", ac.Login)
	//g.POST("/users/logout", ac.Logout)
	//g.POST("/users/token/refresh", ac.RefreshToken)
	//g.GET("/users/token/verify", ac.VerifyToken)
}

func (ctr *auth) Login(c echo.Context) error {
	var credentials *serializers.LoginRequest
	var response *serializers.LoginResponse
	var err error

	if err = c.Bind(&credentials); err != nil {
		bodyErr := errors.NewBadRequestError("failed to parse request body")
		return c.JSON(bodyErr.Status, bodyErr)
	}

	if response, err = ctr.authService.Login(credentials); err != nil {
		switch err {
		case errors.ErrInvalidEmail, errors.ErrInvalidPassword, errors.ErrNotAdmin:
			unAuthErr := errors.NewUnauthorizedError("invalid username or password")
			return c.JSON(unAuthErr.Status, unAuthErr)
		case errors.ErrCreateJwt:
			serverErr := errors.NewInternalServerError("failed to create jwt token")
			return c.JSON(serverErr.Status, serverErr)
		case errors.ErrStoreTokenUuid:
			serverErr := errors.NewInternalServerError("failed to store jwt token uuid")
			return c.JSON(serverErr.Status, serverErr)
		default:
			serverErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
			return c.JSON(serverErr.Status, serverErr)
		}
	}

	return c.JSON(http.StatusCreated, response)
}
