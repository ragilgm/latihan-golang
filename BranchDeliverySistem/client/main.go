package main

import (
	"context"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	BranchDeliverySystem "github.com/ragilmaulana/restapi/tugas-golang/BranchDeliverySistem/bdsProto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:1010"
)

//

func main() {

	// fitur teller ==============================================================
	//_,id,_ := Login(1, "maulana")
	//check := SetorTunai(5555,1000,"kdfkdsmfk",1)
	//fmt.Print(check)
	//over,check := Overbooking(1,4444,5555,100000,"test")
	//fmt.Println(over)
	//fmt.Println(check)
	//check := TarikTunai(5555, 100000, "kdfkdsmfk", 1)
	//fmt.Print(check)
	//checkLogin,cabang,role := Login("ragil", "maulana")
	//cetak,_ := CetakBuku(4444)
	//for _ , value := range cetak.TRANSAKSI{
    //    fmt.Println(value.ID)
	//	fmt.Println(value.NO_REKENING)
	//	fmt.Println(value.JENIS_TRANSAKSI)
	//	fmt.Println(value.TANGGAL)
	//	fmt.Println(value.NOMINAL)
	//	fmt.Println(value.SALDO)
	//	fmt.Println(value.BERITA)
	//}
	//fmt.Println(checkLogin)
	//fmt.Println(cabang)
	//fmt.Println(role)
//========== end fitur teller ================================

// fitur cs =================================================
cif, _ := FindCif(6666)
fmt.Println(cif)

}

//stor tunai client
//====================================================================================================
func SetorTunai(no_req, nominal int, berita string, id_User int) bool {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := BranchDeliverySystem.NewAddClient(conn)
	// call method setor tunai
	respons, err := c.SetorTunai(ctx,
		&BranchDeliverySystem.TRANSAKSI{
			NO_REKENING: int64(no_req),
			NOMINAL:     int64(nominal),
			BERITA:      berita,
			ID:          int64(id_User),
		})

	if respons.Status != 0 {
		fmt.Println(respons)
		return true
	} else {
		return false
	}

}
//====================================================================================================
//tarik tunai client
//====================================================================================================
func TarikTunai(no_req, nominal int, berita string, idUser int) bool {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := BranchDeliverySystem.NewAddClient(conn)
	// call method setor tunai
	respons, err := c.TarikTunai(ctx,
		&BranchDeliverySystem.TRANSAKSI{
			NO_REKENING: int64(no_req),
			NOMINAL:     int64(nominal),
			BERITA:      berita,
			ID:          int64(idUser),
		})

	if respons.Status != 0 {
		fmt.Println(respons)
		return true
	} else {
		return false
	}

}
//====================================================================================================
//Cetak Buku client
//====================================================================================================
func CetakBuku(no_rekening int64) (*BranchDeliverySystem.CETAKBUKU, bool) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := BranchDeliverySystem.NewAddClient(conn)
	// call method setor tunai
	respons, err := c.CetakBuku(ctx,&BranchDeliverySystem.TRANSAKSI{NO_REKENING: no_rekening})
	if respons != nil {
		fmt.Println(respons)
		return respons,true
	} else {
		return respons,false
	}
}
//====================================================================================================
//overbooking client
//====================================================================================================
func Overbooking(id_User int, norekAsal, noRekTujuan, nominal int64, berita string) (*BranchDeliverySystem.OVERBOOKING, bool) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := BranchDeliverySystem.NewAddClient(conn)
	// call method setor tunai
	respons, err := c.OverBooking(ctx,
		&BranchDeliverySystem.OVERBOOKING{
			IdUser: int64(id_User),
			NasabahDetail1: &BranchDeliverySystem.NASABAH_DETAIL{
				NO_REKENING: norekAsal,
			},
			NasabahDetail2: &BranchDeliverySystem.NASABAH_DETAIL{
				NO_REKENING: noRekTujuan,
			},
			Nominal: nominal,
			BERITA:  berita,
		},
	)

	if respons != nil {
		fmt.Println(respons)
		return respons,true
	} else {
		return respons,false
	}
}
//====================================================================================================

// login ==============================================================================================
func Login(id_user int, password string) (bool, string, string) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := BranchDeliverySystem.NewAddClient(conn)

	masuk, err := c.LoginUser(ctx, &BranchDeliverySystem.User{ID_USER: int64(id_user), PASSWORD: password})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	var cabang string
	if masuk.GetNAMA() != "" {
		switch masuk.GetCABANG() {
		case "17600":
			cabang = "jakarta"
			break
		case "17601":
			cabang = "bandung"
			break
		case "17602":
			cabang = "Surabaya"
			break
		}
		return true, cabang, masuk.GetROLE()
	} else {
		return false, cabang, masuk.GetROLE()
	}
}
//=====================================================================================================

// find cif ===========================================================================================
func FindCif(cif int64) (*BranchDeliverySystem.NASABAH, bool) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := BranchDeliverySystem.NewAddClient(conn)
	// call method setor tunai
	respons, err := c.FindCif(ctx,&BranchDeliverySystem.NASABAH{
		CIF: cif,
	})
	if respons.CIF != 0 {
		return respons,true
	} else {
		return respons,false
	}
}
//====================================================================================================
