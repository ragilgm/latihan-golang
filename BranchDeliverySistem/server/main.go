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
func (s *server) SetorTunai(_ context.Context, transaksi *BranchDeliverySystem.TRANSAKSI) (*BranchDeliverySystem.TRANSAKSI, error) {
	db, err := config.GetMysqlDB()
	if err != nil {
		panic(err)
	} else {
		con := services.UserModels{
			DB: db,
		}
		noreq := transaksi.GetNO_REKENING()
		nominal := transaksi.GetNOMINAL()
		berita := transaksi.GetBERITA()

		// respons server
		log.Printf(" client request : %v,%v,%v", noreq, nominal, berita)

		// call method stor tunai for check no rek exist or not
		userDetails, err := con.FindNoRek(int(noreq))
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
func (s *server) TarikTunai(_ context.Context, transaksi *BranchDeliverySystem.TRANSAKSI) (*BranchDeliverySystem.TRANSAKSI, error) {
	db, err := config.GetMysqlDB()
	if err != nil {
		panic(err)
	} else {
		con := services.UserModels{
			DB: db,
		}
		noreq := transaksi.GetNO_REKENING()
		nominal := transaksi.GetNOMINAL()
		berita := transaksi.GetBERITA()

		// respons server
		log.Printf(" client request : %v,%v,%v", noreq, nominal, berita)

		// call method stor tunai for check no rek exist or not
		userDetails, err := con.FindNoRek(int(noreq))
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

// method OverBooking server
func (s *server) OverBooking(_ context.Context, overbooking *BranchDeliverySystem.OVERBOOKING) (*BranchDeliverySystem.OVERBOOKING, error) {
	db, err := config.GetMysqlDB()
	if err != nil {
		panic(err)
	} else {
		con := services.UserModels{
			DB: db,
		}
		rekAwal := overbooking.NasabahDetail1.GetNO_REKENING()
		RekTujuan := overbooking.NasabahDetail2.GetNO_REKENING()
		nominal := overbooking.GetNominal()

		// respons server
		log.Printf(" client request : %v,%v,%v", rekAwal, RekTujuan, nominal)

		// call method stor tunai for check no rek exist or not
		checkReqAwal, err := con.FindNoRek(int(rekAwal))
		fmt.Println(checkReqAwal)
		fmt.Println("called")
		if err != nil {
			panic(err)
		} else {
			fmt.Println("called")
			checkReqTujuan, err := con.FindNoRek(int(RekTujuan))
			fmt.Println(checkReqTujuan)
			if err != nil {
				panic(err)
			} else {
				fmt.Println("called")
				over, err := con.Overbooking(int(overbooking.GetIdUser()),checkReqAwal,checkReqTujuan,int(overbooking.GetNominal()),overbooking.GetBERITA())
				if err != nil {
					panic(err)
				}
				fmt.Println(over)
			}
		}
		return &BranchDeliverySystem.OVERBOOKING{
			NasabahDetail1: &BranchDeliverySystem.NASABAH_DETAIL{
				CIF:         int64(checkReqAwal.CIF),
				NO_REKENING: int64(checkReqAwal.No_Req),
				SALDO:       int64(checkReqAwal.Saldo) - overbooking.GetNominal(),
			},
			NasabahDetail2: &BranchDeliverySystem.NASABAH_DETAIL{
				CIF:         int64(checkReqAwal.CIF),
				NO_REKENING: int64(checkReqAwal.No_Req),
				SALDO:       int64(checkReqAwal.Saldo) - overbooking.GetNominal(),
			},
		}, nil

	}
}

// method cetak buku
func (s *server) CetakBuku(_ context.Context, transaksi *BranchDeliverySystem.TRANSAKSI) (*BranchDeliverySystem.CETAKBUKU, error) {
	db, err := config.GetMysqlDB()
	var listTransaksi []*BranchDeliverySystem.TRANSAKSI
	if err != nil {
		panic(err)
	} else {
		con := services.UserModels{
			db,
		}
		no_rek := transaksi.GetNO_REKENING()

		users, err := con.CetakBuku(int(no_rek))
		if err != nil {
			panic(err)
		}

		for _, velue := range users {
			trx := BranchDeliverySystem.TRANSAKSI{
				ID: int64(velue.Id_Transaksi),
				ID_USER: int64(velue.Id_User),
				NO_REKENING: int64(velue.No_Rekening),
				TANGGAL: velue.Tanggal,
				NOMINAL: int64(velue.Nominal),
				SALDO: int64(velue.Saldo),
				JENIS_TRANSAKSI: velue.Jenis_Transaksi,
				BERITA: velue.Berita,
			}
			listTransaksi = append(listTransaksi,&trx)
			}
		}
	return &BranchDeliverySystem.CETAKBUKU{
		TRANSAKSI: listTransaksi,
	}, nil

}

// method login server
func (s *server) LoginUser(_ context.Context, user *BranchDeliverySystem.User) (*BranchDeliverySystem.User, error) {
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
