package main

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"github.com/ragilmaulana/Latihan/clean/pkg/entity"
	person "github.com/ragilmaulana/Latihan/clean/pkg/proto"
	"github.com/ragilmaulana/Latihan/clean/pkg/usecase"
	"google.golang.org/grpc"
	"log"
	"net"
)

type PersonService struct {
	res usecase.PersonUsecase
}

func (s PersonService) PrintPerson(_ context.Context, user *person.PersonRequest) (*person.PersonRespond, error) {
	req := entity.Person{
		Id_User:   user.GetIdUser(),
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
	}

	res, err := s.res.PrintPerson(req)
	if err != nil {
		panic(err)
	}

	var out *person.PersonRespond

	err1 := mapstructure.Decode(res, &out)
	if err1 != nil {
		return nil,err1
	}
	return out, nil
}

func main() {
	const (
		port = ":1010"
	)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("server starting in port 1010")
	s := grpc.NewServer()
	person.RegisterPersonServiceServer(s, &PersonService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
