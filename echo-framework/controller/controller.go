package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ragilmaulana/bootcamp/tugas-golang/echo-framework/service"
)

func ListUser(c *gin.Context) {
	var user []service.User
	err := service.GetAllUsers(&user)
	if err != nil {
		panic(err)
	}
}

func AddNewUser(c *gin.Context) {
	var user service.User
	c.BindJSON(&user)
	err := service.AddNewUser(&user)
	if err != nil {
	panic(err)
	}
}

func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var user service.User
	err := service.GetSingleUser(&user, id)
	if err != nil {
	panic(err)
	}
}

func PutOneBook(c *gin.Context) {
	var user service.User
	id := c.Params.ByName("id")
	err := service.GetSingleUser(&user, id)
	if err != nil {
		panic(err)
	}
	c.BindJSON(&user)
	err = service.PutOneUser(&user, id)
	if err != nil {
	panic(err)
	}
}

func DeleteBook(c *gin.Context) {
	var user service.User
	id := c.Params.ByName("id")
	err := service.DeleteUser(&user, id)
	if err != nil {
	}
}
