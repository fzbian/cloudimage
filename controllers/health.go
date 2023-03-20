package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Health(c echo.Context) error {
	err := c.String(http.StatusOK, "Ok!")
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
