package main

import (
	"context"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	BranchDeliverySystem "github.com/ragilmaulana/restapi/tugas-golang/BranchDeliverySistem/bdsProto"
	"github.com/ragilmaulana/restapi/tugas-golang/BranchDeliverySistem/config"
	"github.com/ragilmaulana/restapi/tugas-golang/BranchDeliverySistem/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

// method setor tunai server
func (s *server) SetorTunai(ctx context.Context, transaksi *BranchDeliverySystem.TRANSAKSI) (*BranchDeliverySystem.TRANSAKSI, error) {
	db, err := config.GetMysqlDB()
	if err != nil {
		panic(err)
	} else {
		con := services.UserModels{
			db,
		}
		noreq := transaksi.GetNO_REKENING()
		nominal := transaksi.GetNOMINAL()
		berita := transaksi.GetBERITA()

		// respons server
		log.Printf(" client request : %v,%v,%v", noreq, nominal, berita)

		// call method stor tunai for check no rek exist or not
		userDetails, err := con.FindNoRek(int(noreq), int(nominal), berita)
		fmt.Println(userDetails)
		if err != nil {
			panic(err)
		}
		// if no req exist
		if userDetails.No_Req != 0 {
			userId := int(transaksi.GetID_USER())
			trxNominal := int(transaksi.GetNOMINAL())

			// method add setor tunai to db called
			storTunai, err := con.SetorTunaiService(userId, userDetails, transaksi.GetBERITA(), trxNominal)
			if err != nil {
				panic(err)
			}
			if storTunai > 0 {
				fmt.Println("transaksi berhasil")
			} else {
				fmt.Println("Transaksi gagal")
			}
		}
	}

	return &BranchDeliverySystem.TRANSAKSI{}, nil
}

// method tarik tunai server
func (s *server) TarikTunai(ctx context.Context, transaksi *BranchDeliverySystem.TRANSAKSI) (*BranchDeliverySystem.TRANSAKSI, error) {
	db, err := config.GetMysqlDB()
	if err != nil {
		panic(err)
	} else {
		con := services.UserModels{
			db,
		}
		noreq := transaksi.GetNO_REKENING()
		nominal := transaksi.GetNOMINAL()
		berita := transaksi.GetBERITA()

		// respons server
		log.Printf(" client request : %v,%v,%v", noreq, nominal, berita)

		// call method stor tunai for check no rek exist or not
		userDetails, err := con.FindNoRek(int(noreq), int(nominal), berita)
		fmt.Println(userDetails)
		if err != nil {
			panic(err)
		}
		// if no req exist
		if userDetails.No_Req != 0 {
			userId := int(transaksi.GetID_USER())
			trxNominal := int(transaksi.GetNOMINAL())

			// method add tarik tunai to db called
			storTunai, err := con.TarikTunaiService(userId, userDetails, transaksi.GetBERITA(), trxNominal)
			if err != nil {
				panic(err)
			}
			if storTunai > 0 {
				fmt.Println("transaksi berhasil")
			} else {
				fmt.Println("Transaksi gagal")
			}
		}
	}

	return &BranchDeliverySystem.TRANSAKSI{}, nil
}


// method login server
func (s *server) LoginUser(ctx context.Context, user *BranchDeliverySystem.User) (*BranchDeliverySystem.User, error) {
	db, err := config.GetMysqlDB()

	if err != nil {
		panic(err)
	} else {
		con := services.UserModels{
			db,
		}
		id_user := user.GetID_USER()
		password := user.GetPASSWORD()

		users, err := con.Login(int(id_user), password)
		if err != nil {
			panic(err)
		}
		fmt.Println(users)
		u := BranchDeliverySystem.User{
			ID_USER:  int64(users.Id_user),
			NAMA:     users.Nama,
			PASSWORD: users.Password,
			CABANG:   users.Cabang,
			ROLE:     users.Role,
		}
		return &u, nil
	}
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
	BranchDeliverySystem.RegisterAddServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
