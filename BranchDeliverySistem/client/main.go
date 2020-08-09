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

	//check := SetorTunai(4444,1000,"kdfkdsmfk",1)
	//fmt.Print(check)


	check := TarikTunai(4444,100000,"kdfkdsmfk",1)
	fmt.Print(check)
	//checkLogin,cabang,role := Login("ragil", "maulana")
	//fmt.Println(checkLogin)
	//fmt.Println(cabang)
	//fmt.Println(role)

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
	respons , err := c.SetorTunai(ctx,
		&BranchDeliverySystem.TRANSAKSI{
		NO_REKENING: int64(no_req),
		NOMINAL: int64(nominal),
		BERITA: berita,
		ID: int64(id_User),
		})

	if respons != nil {
		fmt.Println(respons)
		return true
	}else {
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
	respons , err := c.TarikTunai(ctx,
		&BranchDeliverySystem.TRANSAKSI{
			NO_REKENING: int64(no_req),
			NOMINAL: int64(nominal),
			BERITA: berita,
			ID: int64(idUser),
		})

	if respons != nil {
		fmt.Println(respons)
		return true
	}else {
		return false
	}

}
//====================================================================================================

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
