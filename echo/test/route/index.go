package route

import (
	"github.com/labstack/echo"
	"net/http"
)

// Index func
func Index() *echo.Echo{
	e := echo.New()

	e.GET("/",func (c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	g := e.Group("/users")
	g.GET("/",)
	return e
}
