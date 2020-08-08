package main

import (
	"bufio"
	"fmt"
	"github.com/chilts/sid"
	"os"
	"time"
)

//==================================================================================
// test 1
//type Test interface {
//	sayHello()
//	sayHay()
//}
//
//type Greeting struct {
//	firstName, lastName string
//}
//
//func (g Greeting) sayHello() (string,string){
//	return g.firstName,g.lastName
//}
//func(g Greeting) sayHai() string{
//	return g.firstName
//}
//
//func main(){
//	g := Greeting{"ragil","maulana"}
//	fmt.Println(g.sayHai()) // ragil
//	fmt.Println(g.sayHello()) // ragil maulana
//}
//==================================================================================

type Parkir struct {
	id            string
	tanggalMasuk  string
	tanggalKeluar string
	platNo        string
	tipekendaraan string
	tarif         int
}

var parkir []Parkir

func (p Parkir) masuk() (string, string) {
	return p.id, p.tanggalMasuk
}

func (p Parkir) keluar() (string, string, string, string, string, int) {
	return p.id, p.tanggalMasuk, p.tanggalKeluar, p.platNo, p.tipekendaraan, p.tarif
}

func findDataMasuk(id string) (bool, string, string) {
	checkData := false
	masuk := ""
	for _, value := range parkir {
		if value.id == id {
			id = value.id
			masuk = value.tanggalMasuk
			checkData = true

		} else {
			checkData = false
		}
	}

	return checkData, id, masuk
}

// hitung tarif
func HitungTarif(in, end time.Time, tipe string) int {
	start := in
	var tarif int = 0
	var tarifLanjutan = 0
	var counter int = 1
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
func GetHargaTarif(tipe string) int {
	var tarif int = 0
	if tipe == "mobil" {
		tarif = 5000
	} else if tipe == "motor" {
		tarif = 3000
	}
	return tarif
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// menu

	inputMenu := ""

	for inputMenu != "99" {
		fmt.Println("SECURE PARKING RAGIL GARAGE \n")

		fmt.Println("1. Parkir Masuk")
		fmt.Println("2. Input Keluar Parkir")
		fmt.Println("99. exit")

		fmt.Print("input pilihan :")
		scanner.Scan()
		inputMenu = scanner.Text()
		switch inputMenu {
		case "1":
			fmt.Println("Generating ID :")
			generateId := sid.Id()
			waktu := time.Now().Format("2006-01-02 15:04:05")
			p := Parkir{id: generateId, tanggalMasuk: waktu}
			i, w := p.masuk()
			parkir = append(parkir, p)
			fmt.Printf("Id Parkir :%v", i)
			fmt.Println()
			fmt.Printf("Tgl Masuk :%v", w)
			fmt.Println()
			break
		case "2":
			fmt.Println("Input ID : ")
			scanner.Scan()
			id := scanner.Text()
			checkId, _, tglMasuk := findDataMasuk(id)
			if checkId {
				fmt.Println("Input Plat No : ")
				scanner.Scan()
				plat := scanner.Text()
				fmt.Println("Input Tipe Kendaraan : ")
				scanner.Scan()
				tipe := scanner.Text()
				waktu := time.Now().Format("2006-01-02 15:04:05")
				parse1, _ := time.Parse("2006-01-02 15:04:05", tglMasuk)
				parse2, _ := time.Parse("2006-01-02 15:04:05", waktu)

				tarif := HitungTarif(parse1, parse2, tipe)
				p := Parkir{id: id, tanggalMasuk: tglMasuk, tanggalKeluar: waktu, platNo: plat, tipekendaraan: tipe, tarif: tarif}
				p.keluar()
				fmt.Printf(" Id Parkir : %v ", p.id)
				fmt.Println()
				fmt.Printf(" Tgl Masuk : %v ", p.tanggalMasuk)
				fmt.Println()
				fmt.Printf(" Tgl Masuk : %v ", p.tanggalKeluar)
				fmt.Println()
				fmt.Printf(" Tgl Masuk : %v ", p.platNo)
				fmt.Println()
				fmt.Printf(" Tgl Masuk : %v ", p.tipekendaraan)
				fmt.Println()
				fmt.Printf(" Tgl Masuk : %v ", p.tarif)
				fmt.Println()
			} else {
				fmt.Println("ID NOT FOUND..")
				break
			}
			break
		default:
			fmt.Println("Something when wrong..!!")
		}
	}
}
