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

	//check := SetorTunai(4444,550,"kdfkdsmfk")
	//fmt.Print(check)
	checkLogin,cabang,role := Login("ragil", "maulana")
	fmt.Println(checkLogin)
	fmt.Println(cabang)
	fmt.Println(role)

}

func SetorTunai(no_req, nominal int, berita string) bool {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := BranchDeliverySystem.NewAddClient(conn)

	respons , err := c.SetorTunai(ctx,
		&BranchDeliverySystem.SETORTUNAI{NOREK: int64(no_req), NOMINAL: int64(nominal), BERITA: berita,
		})

	if respons != nil {
		fmt.Println(respons)
		return true
	}else {
		return false
	}

}


func Login(nama, password string) (bool, string, string) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := BranchDeliverySystem.NewAddClient(conn)

	masuk, err := c.LoginUser(ctx, &BranchDeliverySystem.User{NAMA: "ragil", PASSWORD: "ragil"})
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
