package main

import (
	"context"
	"fmt"
	person "github.com/ragilmaulana/Latihan/clean/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:1010"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := person.NewPersonServiceClient(conn)

	data := person.PersonRequest{
		IdUser: 123,
		FirstName: "test",
		LastName: "ting",
	}
	response, err := c.PrintPerson(ctx, &data)
	fmt.Println("response from server ",response)

}
