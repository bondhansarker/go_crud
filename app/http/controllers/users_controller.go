package controllers

import (
	"demo/app/domain"
	"demo/app/serializers"
	"demo/app/service"
	"demo/app/utils"
	"demo/infra/errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type users struct {
	userService service.IUsers
}

func NewUsersController(e *echo.Echo, userService service.IUsers) {
	userController := &users{
		userService: userService,
	}

	e.GET("/users", userController.Index)
	e.POST("/users", userController.Create)
	e.PUT("/users/:id", userController.Update)
	e.DELETE("/users/:id", userController.Delete)
	e.GET("/users/:id", userController.Show)
}

func (ctr *users) Index(c echo.Context) error {
	result, getErr := ctr.userService.GetUsers()
	if getErr != nil {
		return c.JSON(getErr.Status, getErr)
	}
	return c.JSON(http.StatusOK, result)
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

func (ctr *users) Update(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		return c.JSON(restErr.Status, restErr)
	}

	var user serializers.UserRequest
	if err := c.Bind(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		return c.JSON(restErr.Status, restErr)
	}

	result, updateErr := ctr.userService.UpdateUser(userID, user)

	if updateErr != nil {
		return c.JSON(updateErr.Status, updateErr)
	}

	var response serializers.UserResponse
	respErr := utils.StructToStruct(result, &response)

	if respErr != nil {
		return respErr
	}
	return c.JSON(http.StatusOK, response)
}

func (ctr *users) Show(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		return c.JSON(restErr.Status, restErr)
	}

	result, getErr := ctr.userService.GetUserByID(userID)
	if getErr != nil {
		return c.JSON(getErr.Status, getErr)
	}

	var response serializers.UserResponse
	respErr := utils.StructToStruct(result, &response)

	if respErr != nil {
		return respErr
	}

	return c.JSON(http.StatusOK, response)
}

func (ctr *users) Delete(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		return c.JSON(restErr.Status, restErr)
	}

	getErr := ctr.userService.DeleteUser(userID)
	if getErr != nil {
		return c.JSON(getErr.Status, getErr)
	}
	return c.JSON(http.StatusNoContent, nil)
}
