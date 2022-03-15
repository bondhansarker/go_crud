package controllers

import (
	"fmt"

	"demo/app/serializers"
	"demo/infra/errors"
	"github.com/labstack/echo/v4"
)

func GetUserFromContext(c echo.Context) (*serializers.LoggedInUser, error) {
	user, ok := c.Get("user").(*serializers.LoggedInUser)
	fmt.Println("==============================================")
	fmt.Println(user)
	fmt.Println("==============================================")
	if !ok {
		return nil, errors.ErrNoContextUser
	}

	return user, nil
}
