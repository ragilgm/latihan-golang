package main

import (
	"context"
	"fmt"
	parkir "github.com/ragilmaulana/restapi/tugas-golang/day4/ParkirProto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type ParkirServer struct {
	parkir.UnimplementedParkirAreaServer
}

// get time original
var arrayMasuk []parkir.Masuk
var _ []parkir.Keluar
var _ []parkir.Bill
//FindId(context.Context, *Masuk) (*Validation, error)
func (p *ParkirServer) MasukParkir(_ context.Context, masuk *parkir.Masuk) (*parkir.Masuk, error) {
	log.Println("===========================")
	log.Printf("Parkir Masuk  ")
	log.Printf("Id Parkir : %v ", masuk.GetId())
	log.Printf("Tgl Masuk : %v ", masuk.GetTanggalmasuk())
	log.Println("===========================")

	arrayMasuk = append(arrayMasuk, *masuk)

	return &parkir.Masuk{
		Id:           masuk.GetId(),
		Tanggalmasuk: masuk.GetTanggalmasuk(),
	}, nil
}

func (p *ParkirServer) KeluarParkir(_ context.Context, keluar *parkir.Keluar) (*parkir.Bill, error) {
	var tarif int32 = 0
	for _, index := range arrayMasuk {
		if index.Id == keluar.Id {
			log.Println("===============================================")
			log.Printf("Parkir Keluar  ")
			log.Printf("Id Parkir : %v ", keluar.GetId())
			log.Printf("Tgl Masuk : %v ", index.Tanggalmasuk)
			log.Printf("Tgl Keluar : %v ", keluar.GetTanggalKeluar())
			log.Printf("Tgl Flat : %v ", keluar.GetPlatNo())
			log.Printf("Tgl Tipe : %v ", keluar.GetTipe())
			in, _ := time.Parse("2006-01-02 15:04:05", index.Tanggalmasuk)
			out, _ := time.Parse("2006-01-02 15:04:05", keluar.GetTanggalKeluar())
			tarif = HitungTarif(in, out, keluar.GetTipe())

			log.Printf("Tgl Masuk : %v ", tarif)

			log.Println("===============================================")
			break
		} else {
			return nil, nil
			break
		}
	}
	return &parkir.Bill{
		Tarif: tarif,
	}, nil
}

// hitung tarif
func HitungTarif(in, end time.Time, tipe string) int32 {
	start := in
	var tarif int32 = 0
	var tarifLanjutan int32 = 0
	var counter int32 = 1
	affter := start
	for affter != end {
		affter = start.Add(time.Second * time.Duration(counter))
		counter++
	}
	harga := GetHargaTarif(tipe)

	switch tipe {
	case "mobil":
		tarifLanjutan = 3000
		if counter == 1 {
			tarif = harga
		} else {
			tarif = harga + (tarifLanjutan * counter)
		}
		break
	case "motor":
		tarifLanjutan = 2000
		if counter == 1 {
			tarif = harga
		} else {
			tarif = harga + (tarifLanjutan * counter)
		}
		break
	default:
		fmt.Println("wrong input")
		break
	}
	return tarif

}

// get harga tarif
func GetHargaTarif(tipe string) int32 {
	var tarif int32 = 0
	if tipe == "mobil" {
		tarif = 5000
	} else if tipe == "motor" {
		tarif = 3000
	}
	return tarif
}

func (p *ParkirServer) FindId(_ context.Context, check *parkir.Validation) (*parkir.Validation, error) {
	checkId := false
	for _ , value := range arrayMasuk{
		if(value.Id == check.Id ){
         checkId = true
			fmt.Println("id ditemukan")
         break
		}else {
			checkId = false
			fmt.Println(checkId)
			break
		}
	}
	return &parkir.Validation{Check: checkId,Id: check.Id }, nil
}

func main() {
	const (
		port = ":8080"
	)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	parkir.RegisterParkirAreaServer(s, &ParkirServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
