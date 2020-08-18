package route

import (
	"github.com/labstack/echo"
	"github.com/ragilmaulana/restapi/tugas-golang/echo/test/controller"
	"github.com/ragilmaulana/restapi/tugas-golang/echo/test/models"
)

var users = models.Users{
	models.User{},
}

func UserRoute(g *echo.Group) {
	g.POST("/", controller.AddUser)
}
