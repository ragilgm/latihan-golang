package main

import (
	"fmt"
	testProto "github.com/ragilmaulana/restapi/tugas-golang/Latihan/grpcAPI/proto"
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
	get(con, g)
	multiply(con,g)


	if err = g.Run(":8080") ; err != nil {
		log.Fatal("failed to run server :",err)
	}
}


func get(con *grpc.ClientConn, g *gin.Engine){
	client := testProto.NewAddClient(con)

	g.GET("/add/:a/:b", func(ctx *gin.Context) {

		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter B"})
			return
		}

		req := &testProto.Request{A: int64(a), B: int64(b)}
		if respons,err := client.Add(ctx,req); err == nil {
			ctx.JSON(http.StatusOK,gin.H{
				"result" : fmt.Sprint(respons.Result),
			})
		}else {
			ctx.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
		}


	})

}

func multiply (con *grpc.ClientConn,g *gin.Engine){
	client := testProto.NewAddClient(con)
	g.GET("/mutiply/:a/:b", func(ctx *gin.Context) {

		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter B"})
			return
		}

		req := &testProto.Request{A: int64(a), B: int64(b)}
		if respons,err := client.Add(ctx,req); err == nil {
			ctx.JSON(http.StatusOK,gin.H{
				"result" : fmt.Sprint(respons.Result),
			})
		}else {
			ctx.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
		}


	})


}