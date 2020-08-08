package main

import (
	"context"
	testProto "github.com/ragilmaulana/Latihan/grpcAPI/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct {
}

func main() {

	Listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	testProto.RegisterAddServer(s, &server{})
	reflection.Register(s)
if e := s.Serve(Listener); e != nil {
	panic(e)
}
}

func (s *server) Add(ctx context.Context, request *testProto.Request) (*testProto.Responds, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &testProto.Responds{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *testProto.Request) (*testProto.Responds, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &testProto.Responds{Result: result}, nil
}


