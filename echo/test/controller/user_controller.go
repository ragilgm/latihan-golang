package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func AddUser(c echo.Context) error {
	return c.String(http.StatusOK, "group route")
}