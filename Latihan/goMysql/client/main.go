package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	UserService "github.com/ragilmaulana/Latihan/goMysql/protoUser"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

func main() {

	con, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	g := gin.Default()

	// router
	getSingleUser(con, g)
	GetAllUser(con, g)
	UpdateUser(con, g)
	Insert(con,g)

	if err = g.Run(":8080"); err != nil {
		log.Fatal("failed to run server :", err)
	}
}

func getSingleUser(con *grpc.ClientConn, g *gin.Engine) {
	client := UserService.NewAddClient(con)
	g.GET("/user/:id", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter A"})
			return
		}
		if respons, err := client.FindID(ctx, &UserService.User{ID: int64(a)}); err == nil {

			ctx.JSON(http.StatusOK, respons)

		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
}

func GetAllUser(con *grpc.ClientConn, g *gin.Engine) {
	client := UserService.NewAddClient(con)
	g.GET("/users", func(ctx *gin.Context) {

		if respons, err := client.FindAll(ctx, &UserService.Empty{}); err == nil {
			ctx.JSON(http.StatusOK, respons.GetUser())
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
}

func UpdateUser(con *grpc.ClientConn, g *gin.Engine) {
	client := UserService.NewAddClient(con)
	g.PUT("/user/:id", func(ctx *gin.Context) {
		body := ctx.Request.Body
		decoder := json.NewDecoder(body)
		var User UserService.User
		err := decoder.Decode(&User)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad Request"})
			return
		}
		if respons, err := client.EditData(ctx, &UserService.User{ID: User.ID, FIRSTNAME: User.FIRSTNAME, LASTNAME: User.LASTNAME}); err == nil {
			ctx.JSON(http.StatusOK, respons)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
}

func Insert(con *grpc.ClientConn, g *gin.Engine) {
	client := UserService.NewAddClient(con)
	g.POST("/user", func(ctx *gin.Context) {
		body := ctx.Request.Body
		decoder := json.NewDecoder(body)
		var User UserService.User
		err := decoder.Decode(&User)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad Request"})
			return
		}
		if respons, err := client.InsertData(ctx, &UserService.User{ID: User.ID, FIRSTNAME: User.FIRSTNAME, LASTNAME: User.LASTNAME}); err == nil {
			ctx.JSON(http.StatusOK, respons)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
}
