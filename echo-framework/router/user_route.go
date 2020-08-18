package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func GetRouter() {

	server := echo.New()

	server.Use(middleware.Logger())

	server.Use(middleware.Recover())


	server.GET("/", GetUser)




}
func GetUser(c echo.Context) error {
	//render with master
	return c.Render(http.StatusOK, "index", "hello")
}

func getAllUser(c echo.Context) error {
	//render only file, must full name with extension
	return c.Render(http.StatusOK, "page.tpl", echo.Map{"title": "Page file title!!"})
}
