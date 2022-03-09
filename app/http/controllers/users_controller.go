package controllers

import (
	"demo/app/domain"
	"demo/app/serializers"
	"demo/app/service"
	"demo/app/utils"
	"demo/infra/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type users struct {
	userService service.IUsers
}

func NewUsersController(e *echo.Echo, userService service.IUsers) {

	userController := &users{
		userService: userService,
	}

	e.POST("/users/", userController.Create)
	e.GET("/users/", userController.Index)
	// e.GET("/users/{id}", controllers.Show)
	// e.PUT("/users/{id}", controllers.Update)
	// e.DELETE("/users/{id}", controllers.Delete)
}

func (ctr *users) Create(c echo.Context) error {

	var user domain.User

	if err := c.Bind(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		return c.JSON(restErr.Status, restErr)
	}

	result, saveErr := ctr.userService.CreateUser(user)
	if saveErr != nil {
		return c.JSON(saveErr.Status, saveErr)
	}
	var response serializers.UserResponse
	respErr := utils.StructToStruct(result, &response)
	if respErr != nil {
		return respErr
	}

	return c.JSON(http.StatusCreated, response)
}
func (ctr *users) Index(c echo.Context) error {

	result, getErr := ctr.userService.GetUsers()
	if getErr != nil {
		return c.JSON(getErr.Status, getErr)
	}

	return c.JSON(http.StatusCreated, result)
}
