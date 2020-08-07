package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Masuk struct {
	IDPARKIR     string    `json:"indParking"`
	TANGGALMASUK time.Time `json:"tanggalMasuk"`
}

type Keluar struct {
	TANGGALKELUAR time.Time `json:"tanggalMasuk"`
	PLAT          string    `json:"plat"`
	TIPE          string    `json:"tipe"`
}

type Parking struct {
	MASUK  Masuk
	KELUAR Keluar
	TARIF  int
}
var arrayMasuk []Masuk

// set masuk
func getDataMasuk() {
	resp, err := http.Get("http://localhost:8080/api/list/in")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//var masuk Masuk
	var jsonR = arrayMasuk
	err = json.Unmarshal([]byte(body), &jsonR)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("\nDATA MASUK : ")
	for _, value := range jsonR {
		fmt.Println("ID \t: ", value.IDPARKIR)
		fmt.Println("JAM MASUK : ", value.TANGGALMASUK)
		arrayMasuk = append(arrayMasuk,value)
	}
	fmt.Println(arrayMasuk)

}

// find id
func findId(id string) bool {
	resp, err := http.Get("http://localhost:8080/api/list/" + id + "")
	if err != nil {
		// handle error
		fmt.Println(resp)
		fmt.Println(id)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//var masuk Masuk
	var jsonR = Masuk{}
	err = json.Unmarshal(body, &jsonR)
	if err != nil {
		log.Println(err)
	}
	check := false
	if jsonR.IDPARKIR == id {
		check = true
	} else {
		check = false
	}

	return check
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

func printTagihan(p Parking) {
	fmt.Println("ID : ", p.MASUK.IDPARKIR)
	fmt.Println("TANGGAL MASUK : ", p.MASUK.TANGGALMASUK)
	fmt.Println("TANGGAL KELUAR : ", p.KELUAR.TANGGALKELUAR)
	fmt.Println("TIPE KENDARAAN : ", p.KELUAR.TIPE)
	fmt.Println("FLAT KENDARAAN : ", p.KELUAR.PLAT)
	p.TARIF = HitungTarif(p.MASUK.TANGGALMASUK, p.KELUAR.TANGGALKELUAR, p.KELUAR.TIPE)
	fmt.Println("TARIF \t : ", p.TARIF)
}

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

func main() {

	// scanner
	scanner := bufio.NewScanner(os.Stdin)

	// menu
	var input string = ""
	for input != "99" {
		fmt.Println("SECURE PARKING RAGIL GARAGE \n")

		fmt.Println("1. Print status")
		fmt.Println("2. Input Keluar Parkir")
		fmt.Println("99. exit")

		fmt.Print("input pilihan :")
		scanner.Scan()
		input = scanner.Text()
		switch input {
		case "1":
			getDataMasuk()
			fmt.Println()
			//getDataKeluar()
			break
		case "2":
			fmt.Println("Input ID : ")
			scanner.Scan()
			id := scanner.Text()
			checkId := findId(id)
			if checkId {
				fmt.Println("Input Plat No : ")
				scanner.Scan()
				plat := scanner.Text()
				fmt.Println("Input Tipe Kendaraan : ")
				scanner.Scan()
				tipe := scanner.Text()
				var parking Parking
				parking.KELUAR.TANGGALKELUAR = GetTime()
				parking.KELUAR.TIPE = tipe
				parking.KELUAR.PLAT = plat
				for _, value := range arrayMasuk {
					if value.IDPARKIR == id {
						parking.MASUK.IDPARKIR = value.IDPARKIR
						parking.MASUK.TANGGALMASUK = value.TANGGALMASUK
					}
				}
				printTagihan(parking)
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
