package main

import (
	"context"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	BranchDeliverySystem "github.com/ragilmaulana/restapi/tugas-golang/BranchDeliverySistem/bdsProto"
	"github.com/ragilmaulana/restapi/tugas-golang/BranchDeliverySistem/config"
	"github.com/ragilmaulana/restapi/tugas-golang/BranchDeliverySistem/entities"
	"github.com/ragilmaulana/restapi/tugas-golang/BranchDeliverySistem/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

// method setor tunai server
func (s *server) SetorTunai(_ context.Context, transaksi *BranchDeliverySystem.TRANSAKSI) (*BranchDeliverySystem.STATUS, error) {
	var status int
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
			var pointerStatus *int = &status
			if storTunai > 0 {
				fmt.Println("transaksi berhasil")
				*pointerStatus = 1
				return &BranchDeliverySystem.STATUS{Status: int64(*pointerStatus)}, nil
			} else {
				*pointerStatus = 0
				fmt.Println("transaksi gagal")
				return &BranchDeliverySystem.STATUS{Status: int64(*pointerStatus)}, nil
			}
		}
	}
	return &BranchDeliverySystem.STATUS{
		Status: int64(status),
	}, nil

}

// method tarik tunai server ========================================================================================================
func (s *server) TarikTunai(_ context.Context, transaksi *BranchDeliverySystem.TRANSAKSI) (*BranchDeliverySystem.STATUS, error) {
	var status int
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
			var pointerStatus *int = &status
			if storTunai > 0 {
				*pointerStatus = 1
				fmt.Println("transaksi berhasil")
			} else {
				*pointerStatus = 0
				fmt.Println("Transaksi gagal")
			}
		}
	}

	return &BranchDeliverySystem.STATUS{
		Status: int64(status),
	}, nil
}

// end of method tarik tunai server =================================================================================================

// method OverBooking server ========================================================================================================
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
		}
			fmt.Println("called")
			checkReqTujuan, err := con.FindNoRek(int(RekTujuan))
			fmt.Println(checkReqTujuan)
			if err != nil {
				panic(err)
			} else {
				fmt.Println("called")
				over, err := con.Overbooking(int(overbooking.GetIdUser()), checkReqAwal, checkReqTujuan, int(overbooking.GetNominal()), overbooking.GetBERITA())
				if err != nil {
					panic(err)
				}
				fmt.Println(over)
			}

		return &BranchDeliverySystem.OVERBOOKING{
			NasabahDetail1: &BranchDeliverySystem.NASABAH_DETAIL{
				CIF:         int64(checkReqAwal.CIF),
				NO_REKENING: int64(checkReqAwal.No_Req),
				SALDO:       int64(checkReqAwal.Saldo) - overbooking.GetNominal(),
			},
			NasabahDetail2: &BranchDeliverySystem.NASABAH_DETAIL{
				CIF:         int64(checkReqTujuan.CIF),
				NO_REKENING: int64(checkReqTujuan.No_Req),
				SALDO:       int64(checkReqTujuan.Saldo) + overbooking.GetNominal(),
			},
		}, nil

	}
}

// end of method OverBooking server =================================================================================================

// method cetak buku ================================================================================================================
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
				ID:              int64(velue.Id_Transaksi),
				ID_USER:         int64(velue.Id_User),
				NO_REKENING:     int64(velue.No_Rekening),
				TANGGAL:         velue.Tanggal,
				NOMINAL:         int64(velue.Nominal),
				SALDO:           int64(velue.Saldo),
				JENIS_TRANSAKSI: velue.Jenis_Transaksi,
				BERITA:          velue.Berita,
			}
			listTransaksi = append(listTransaksi, &trx)
		}
	}
	return &BranchDeliverySystem.CETAKBUKU{
		TRANSAKSI: listTransaksi,
	}, nil

}

// end of method cetak buku ========================================================================================================

// method Find Nasabah by nik ========================================================================================================
func (s *server) FindByNIK(_ context.Context, nasabah *BranchDeliverySystem.NASABAH) (*BranchDeliverySystem.NASABAH, error) {
	db, err := config.GetMysqlDB()

	con := services.UserModels{
		db,
	}
	nik := nasabah.GetNIK()
	fmt.Println(nik)
	n, err := con.FindNik(int(nik))
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	u := BranchDeliverySystem.NASABAH{
		CIF:           int64(n.CIF),
		NIK:           int64(n.NIK),
		NAMA:          n.Nama,
		TEMPAT_LAHIR:  n.Tempat_Lahir,
		TANGGAL_LAHIR: n.Tanggal_Lahir,
		ALAMAT:        n.Alamat,
		NO_TELP:       nasabah.NO_TELP,
	}
	return &u, nil
}

//=============== end of find nik ===================================================================================================

// method Find Nasabah by Cif ========================================================================================================
func (s *server) FindByCIF(_ context.Context, nasabah *BranchDeliverySystem.NASABAH) (*BranchDeliverySystem.NASABAH_INFO, error) {
	db, err := config.GetMysqlDB()

	con := services.UserModels{
		db,
	}
	cif := nasabah.GetCIF()
	fmt.Println(cif)
	n, err := con.PrintNasabahInfoByCif(int(cif))
	if err != nil {
		panic(err)
	}

	if n.NasabahDetail.CIF == 0 {
		return nil, nil
	} else {
		return &BranchDeliverySystem.NASABAH_INFO{
			NASABAH: &BranchDeliverySystem.NASABAH{
				CIF:           int64(n.Nasabah.CIF),
				NIK:           int64(n.Nasabah.NIK),
				NAMA:          n.Nasabah.Nama,
				TEMPAT_LAHIR:  n.Nasabah.Tempat_Lahir,
				TANGGAL_LAHIR: n.Nasabah.Tanggal_Lahir,
				ALAMAT:        n.Nasabah.Alamat,
				NO_TELP:       n.Nasabah.No_Telp,
			}, NASABAH_DETAIL: &BranchDeliverySystem.NASABAH_DETAIL{
				CIF:         int64(n.NasabahDetail.CIF),
				NO_REKENING: int64(n.NasabahDetail.No_Req),
				SALDO:       int64(n.NasabahDetail.Saldo),
			},
		}, nil
	}
}

//=============== end of find cif ===================================================================================================

// Buat cif ========================================================================================================
func (s *server) BuatCif(_ context.Context, nasabah *BranchDeliverySystem.NASABAH) (*BranchDeliverySystem.NASABAH, error) {
	db, err := config.GetMysqlDB()

	con := services.UserModels{
		db,
	}
	create := entities.Nasabah{
		CIF:           int(nasabah.GetCIF()),
		NIK:           int(nasabah.GetNIK()),
		Nama:          nasabah.GetNAMA(),
		Tempat_Lahir:  nasabah.GetTEMPAT_LAHIR(),
		Tanggal_Lahir: nasabah.GetTANGGAL_LAHIR(),
		Alamat:        nasabah.GetALAMAT(),
		No_Telp:       nasabah.GetNO_TELP(),
	}
	n, err := con.BuatCIF(create)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	return &BranchDeliverySystem.NASABAH{
		CIF:           int64(n.CIF),
		NIK:           int64(n.NIK),
		NAMA:          n.Nama,
		TEMPAT_LAHIR:  n.Tempat_Lahir,
		TANGGAL_LAHIR: n.Tanggal_Lahir,
		ALAMAT:        n.Alamat,
		NO_TELP:       n.No_Telp,
	}, nil
}

//=============== end of buat ===================================================================================================

// Buat tabungan ========================================================================================================
func (s *server) BuatTabungan(_ context.Context, nasabah *BranchDeliverySystem.NASABAH_INFO) (*BranchDeliverySystem.NASABAH_INFO, error) {
	db, err := config.GetMysqlDB()

	con := services.UserModels{
		db,
	}
	cif := nasabah.NASABAH.GetCIF()
	saldo := nasabah.NASABAH_DETAIL.GetSALDO()

	n, err := con.BuatTabungan(cif, saldo)
	if err != nil {
		panic(err)
	}
	return &BranchDeliverySystem.NASABAH_INFO{
		NASABAH: &BranchDeliverySystem.NASABAH{
			CIF:           int64(n.Nasabah.CIF),
			NIK:           int64(n.Nasabah.NIK),
			NAMA:          n.Nasabah.Nama,
			TEMPAT_LAHIR:  n.Nasabah.Tempat_Lahir,
			TANGGAL_LAHIR: n.Nasabah.Tanggal_Lahir,
			ALAMAT:        n.Nasabah.Alamat,
			NO_TELP:       n.Nasabah.No_Telp,
		}, NASABAH_DETAIL: &BranchDeliverySystem.NASABAH_DETAIL{
			CIF:         int64(n.NasabahDetail.CIF),
			NO_REKENING: int64(n.NasabahDetail.No_Req),
			SALDO:       int64(n.NasabahDetail.Saldo),
		},
	}, nil
}

//=============== end of buat ===================================================================================================

// method login server ==============================================================================================================
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

// end of login user ==============================================================================================================

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
