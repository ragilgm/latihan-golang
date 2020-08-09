package main

import (
	"context"
	"fmt"
	"github.com/ragilmaulana/Latihan/goMysql/config"
	"github.com/ragilmaulana/Latihan/goMysql/entities"
	"github.com/ragilmaulana/Latihan/goMysql/models"
	UserService "github.com/ragilmaulana/Latihan/goMysql/protoUser"
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

	UserService.RegisterAddServer(s, &server{})
	reflection.Register(s)
	if e := s.Serve(Listener); e != nil {
		panic(e)
	}
}

func (s *server) InsertData(_ context.Context, user *UserService.User) (*UserService.User, error) {
	db, err := config.GetMysqlDB()
	if err != nil {
		panic(err)
	} else {
		models := models.UserModels{
			db,
		}
		user := entities.User{
			FIRSTNAME: user.GetFIRSTNAME(),
			LASTNAME:  user.GetLASTNAME(),
		}
		users, err := models.Insert(&user)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Id", users, "have been inserted")

		}
		return &UserService.User{}, nil
	}
}

func (s *server) EditData(_ context.Context, user *UserService.User) (*UserService.User, error) {
	db, err := config.GetMysqlDB()
	if err != nil {
		panic(err)
	} else {
		models := models.UserModels{
			db,
		}
		U := entities.User{
			ID:        int(user.GetID()),
			FIRSTNAME: user.GetFIRSTNAME(),
			LASTNAME:  user.GetLASTNAME(),
		}

		users, err := models.Update(&U)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Id", users, "have been inserted")

		}
		return &UserService.User{}, nil
	}
}

func (s *server) FindAll(_ context.Context, _ *UserService.Empty) (*UserService.Users, error) {
	db, err := config.GetMysqlDB()
	var ListUser []*UserService.User
	if err != nil {
		panic(err)
	} else {
		user := models.UserModels{
			db,
		}
		users, err := user.FindALL()
		if err != nil {
			panic(err)
		}
		for _, velue := range users {
			u := UserService.User{
				ID:        int64(velue.ID),
				FIRSTNAME: velue.FIRSTNAME,
				LASTNAME:  velue.LASTNAME,
			}

			ListUser = append(ListUser, &u)
		}
	}
	return &UserService.Users{User: ListUser}, nil
}

func (s *server) FindID(_ context.Context, user *UserService.User) (*UserService.User, error) {
	db, err := config.GetMysqlDB()
	if err != nil {
		panic(err)
	} else {
		con := models.UserModels{
			db,
		}
		id := int(user.GetID())
		users, err := con.FIndById(id)
		if err != nil {
			panic(err)
		}
		u := UserService.User{
			ID:        int64(users.ID),
			FIRSTNAME: users.FIRSTNAME,
			LASTNAME:  users.LASTNAME,
		}
		return &u, nil
	}
}
