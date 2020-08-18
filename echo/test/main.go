package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/ragilmaulana/restapi/tugas-golang/echo/test/route"
)

func main(){
	e := route.Index()
	e.Use(middleware.Logger())

	e.Start(":9000")
}