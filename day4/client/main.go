package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/chilts/sid"
	parkir "github.com/ragilmaulana/restapi/tugas-golang/day4/ParkirProto"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func GetTime() time.Time {
	//init the loc
	loc, _ := time.LoadLocation("Asia/Jakarta")
	//set timezone,
	currentTime := time.Now().In(loc)

	timeStampString := currentTime.Format("2006-01-02 15:04:05")
	layOut := "2006-01-02 15:04:05"
	timeStamp, err := time.Parse(layOut, timeStampString)
	if err != nil {
		fmt.Println(err)
	}
	return timeStamp
}

func parkirMasuk() *parkir.Masuk {
	idParkir := sid.Id()
	tglMasuk := GetTime().Format("2006-01-02 15:04:05")

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := parkir.NewParkirAreaClient(conn)

	masuk, err := c.MasukParkir(ctx, &parkir.Masuk{Id: idParkir, Tanggalmasuk: tglMasuk})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return masuk
}

func parkirKeluar(id, plat, tipe string) *parkir.Bill {
	idParkir := id
	fmt.Println(idParkir)
	tglKeluar := GetTime().Format("2006-01-02 15:04:05")

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := parkir.NewParkirAreaClient(conn)

	keluar, err := c.KeluarParkir(ctx, &parkir.Keluar{
		Id:            id,
		TanggalKeluar: tglKeluar,
		PlatNo:        plat,
		Tipe:          tipe,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return keluar
}

func FindId(id string) *parkir.Validation {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := parkir.NewParkirAreaClient(conn)

	check, err := c.FindId(ctx, &parkir.Validation{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return check
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// menu
	var input string = ""
	for input != "99" {
		fmt.Println("SECURE PARKING RAGIL GARAGE \n")

		fmt.Println("TARIF\n")
		fmt.Println(" detik pertama untuk motor : 3.000")
		fmt.Println(" detik berikutnya untuk motor : 2000\n")
		fmt.Println(" detik pertama untuk Mobil : 5.000")
		fmt.Println(" detik berikutnya untuk motor : 3000\n")

		fmt.Println("1. Parkir Masuk")
		fmt.Println("1. Keluar Parkir")
		fmt.Println("99. exit")

		fmt.Print("input pilihan :")
		scanner.Scan()
		input = scanner.Text()
		switch input {
		case "1":
			masuk := parkirMasuk()
			log.Println("=================================================")
			log.Printf("Parkir Masuk  ")
			log.Printf("Id Parkir : %v ", masuk.GetId())
			log.Printf("Tgl Masuk : %v ", masuk.GetTanggalmasuk())
			log.Println("=================================================")
			break
		case "2":
			fmt.Println("Masukan ID Kendaraan : ")
			scanner.Scan()
			id := scanner.Text()

			//check id exist or not
			if FindId(id).GetCheck() {

				fmt.Println("Masukan FLAT Motor : ")
				scanner.Scan()
				flat := scanner.Text()
				fmt.Println("Masukan Tipe Kendaraan : ")
				scanner.Scan()
				tipe := scanner.Text()
				keluar := parkirKeluar(id, flat, tipe)
				fmt.Println("=================== Total ==================")
				fmt.Println("============= Rp.",keluar.GetTarif()," ==================")
				fmt.Println("============================================\n")
			} else {
				fmt.Println("id not found")
			}
		case "99":
			fmt.Printf("Logout SIstem")
			break
		default:
			fmt.Println("Something when wrong..!!")
			break
		}

	}

}
