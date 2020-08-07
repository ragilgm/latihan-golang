package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Masuk struct {
	tanggalMasuk time.Time
	indParking   uuid.UUID
	status       Status
}

type Keluar struct {
	tanggalMasuk  time.Time
	tanggalKeluar time.Time
	indParking    uuid.UUID
	Kendaraan     Kendaraan
	status        Status
	tarif         int
}

type Kendaraan struct {
	tipeKendaraan string
	platNo        string
}

type Status struct {
	status string
}

var masuk = make(map[uuid.UUID]Masuk)
var keluar = make(map[uuid.UUID]Keluar)

// generate id
func GenerateId() uuid.UUID {
	u1 := uuid.Must(uuid.NewV4())
	return u1
}

// set masuk
func (m Masuk) setMasuk(tglMasuk time.Time, id uuid.UUID) {
	m.tanggalMasuk = tglMasuk
	m.indParking = id
	m.status.status = "masuk"
	masuk[id] = m
}

// set keluar
func (k Keluar) setKeluar(tglKeluar, tglMasuk time.Time, id uuid.UUID, tipe, plat string, tarif int) {
	k.tanggalMasuk = tglMasuk
	k.tanggalKeluar = tglKeluar
	k.indParking = id
	k.Kendaraan.tipeKendaraan = tipe
	k.Kendaraan.platNo = plat
	k.status.status = "keluar"
	k.tarif = tarif
	keluar[id] = k
}

// get time original
func getTime() time.Time {
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

// check id is exist
func idExist(id uuid.UUID) bool {
	value := false
	for _, index := range masuk {
		if index.indParking == id {
			value = true
			break
		} else {
			value = false
			break
		}
	}
	return value
}

// get harga tarif
func getHargaTarif(tipe string) int {
	var tarif int = 0
	if tipe == "mobil" {
		tarif = 5000
	} else {
		tarif = 3000
	}
	return tarif
}

// hitung tarif
func hitungTarif(in, end time.Time, tipe string) int {
	start := in
	var tarif int = 0
	var tarifLanjutan = 0
	var counter int = 1
	affter := start
	for affter != end {
		affter = start.Add(time.Second * time.Duration(counter))
		counter++
	}
	harga := getHargaTarif(tipe)

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

// remove key dari map masuk jika id sudah keluar
func removeIdMasuk(id uuid.UUID) {
	for key, value := range masuk {
		if value.indParking == id {
			delete(masuk, key)
			break
		} else {
			fmt.Println(("id tidak di temukan"))
			break
		}
	}
	return
}

// print total data masuk
func printMasuk() {
	fmt.Println("DATA KENDARAAN MASUK")
	for _, value := range masuk {
		fmt.Println("id : ", value.indParking, "\ttanggal : ", value.tanggalMasuk.Format("2006-01-02 15:04:05"), "\tstatus : ", value.status.status)
		fmt.Println()
	}
}

// print total data keluar
func printKeluar() {
	fmt.Println("\nDATA KENDARAAN kELUAR")
	for _, value := range keluar {
		fmt.Println("ID Parkir : ", value.indParking)
		fmt.Println("Tipe Kendaraan : ", value.Kendaraan.tipeKendaraan)
		fmt.Println("Plat No : ", value.Kendaraan.platNo)
		fmt.Println("Tanggal Masuk : ", value.tanggalMasuk.Format("2006-01-02 15:04:05"))
		fmt.Println("Tanggal Keluar : ", value.tanggalKeluar.Format("2006-01-02 15:04:05"))
		fmt.Println("Status : ", value.status.status)
		fmt.Println("Tarif : ", value.tarif, "\n")
	}
}

func main() {
	// init map masuk
	var in Masuk

	// init map keluar
	var out Keluar

	// scanner
	scanner := bufio.NewScanner(os.Stdin)

	// menu
	var input string = ""
	for input != "99" {
		start := getTime()
		fmt.Println("SECURE PARKING RAGIL GARAGE \n")

		fmt.Println("TARIF\n")
		fmt.Println("1 detik pertama untuk motor : 3.000")
		fmt.Println("1 detik berikutnya untuk motor : 2000\n")
		fmt.Println("1 detik pertama untuk Mobil : 5.000")
		fmt.Println("1 detik berikutnya untuk motor : 3000\n")

		fmt.Println("1. Parkir Masuk")
		fmt.Println("2. Parkir Keluar")
		fmt.Println("3. Print status")
		fmt.Println("99. exit")

		fmt.Print("input pilihan :")
		scanner.Scan()
		input = scanner.Text()
		switch input {
		case "1":
			id := GenerateId()
			fmt.Println("Jam Masuk \t: ", start.Format("2006-01-02 15:04:05"))
			fmt.Println("ID Tiket \t: ", id)
			fmt.Println("\nSILAHKAN MASUK JURAGAN...!!!\n")
			in.setMasuk(start, id)
			break
		case "2":
			fmt.Println("input Id Tiket : ")
			scanner.Scan()
			id := scanner.Text()
			u2, _ := uuid.FromString(id)

			checkId := idExist(u2)

			if checkId {
				end := getTime()
				fmt.Println("Jam Keluar : ", end.Format("2006-01-02 15:04:05"))
				fmt.Println("input plat nomer : ")
				scanner.Scan()
				plat := scanner.Text()
				fmt.Println("input Tipe Kendaraan : ")
				scanner.Scan()
				tipe := scanner.Text()
				tarif := hitungTarif(start, end, tipe)
				fmt.Println("tarif : ", tarif)
				out.setKeluar(start, end, u2, tipe, plat, tarif)
				removeIdMasuk(u2)
				fmt.Println("\nSELAMAT JALAN JANGAN LUPA BERDO'A YAH...!!!\n")
			} else {
				fmt.Println("id not found")
			}
			break
		case "3":
			printMasuk()
			printKeluar()
			break
		default:
			fmt.Println("Something when wrong..!!")
		}

	}
}
