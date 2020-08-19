package main

import (
	"context"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/ragilmaulana/Latihan/clean/pkg/domain"
	person "github.com/ragilmaulana/Latihan/clean/pkg/proto"
	"github.com/ragilmaulana/Latihan/clean/pkg/usecase"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	input usecase.PersonUsecase
}

func NewPersonService(a usecase.PersonUsecase) *server {
	return &server{
		input: a,
	}
}

func (s server) PrintPerson(ctx context.Context, empty *person.Empty) (*person.AllResopons, error) {
	res, err := s.input.PrintPerson()
	if err != nil {
		panic(err)
	}

	var out []*person.PersonResponse
	err1 := mapstructure.Decode(res, &out)
	if err1 != nil {
		return nil, err1
	}
	fmt.Println(out)
	return &person.AllResopons{
		PersonResponse: out,
	}, nil
}

func (s server) AddPerson(ctx context.Context, request *person.PersonRequest) (*person.Message, error) {
	req := &domain.Person{
		Id_User:   request.GetIdUser(),
		FirstName: request.GetFirstName(),
		LastName:  request.GetLastName(),
	}
	respond, err := s.input.AddPerson(req)
	if err != nil {
		panic(err)
	}
	if respond != 1 {
		return &person.Message{
			Message: "data gagal ditambahkan",
		}, nil
	} else {
		return &person.Message{
			Message: "data berhasil di tambahkan",
		}, nil
	}
}

func (s server) DeletePerson(ctx context.Context, request *person.PersonRequest) (*person.Message, error) {
	req := &domain.Person{
		Id_User: request.GetIdUser(),
	}
	responds, err := s.input.DeletePerson(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(responds)
	if responds != 1 {
		return &person.Message{
			Message: "data gagal dihapus",
		}, nil
	} else {
		return &person.Message{
			Message: "data berhasil di hapus",
		}, nil
	}
}

func (s server) EditPersonById(ctx context.Context, request *person.PersonRequest) (*person.PersonResponse, error) {
	panic("implement me")
}

func (s server) EditPersonByName(ctx context.Context, request *person.PersonRequest) (*person.PersonResponse, error) {
	panic("implement me")
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
	person.RegisterPersonServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
