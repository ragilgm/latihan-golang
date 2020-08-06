package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"os"
	"time"
)

type Parkir struct {
	indParking    uuid.UUID `json:"indParking"`
}


// generate id
func GenerateId() uuid.UUID {
	u1 := uuid.Must(uuid.NewV4())
	return u1
}

// get time original
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

// set masuk
func (p Parkir) SetMasuk() {
	id := GenerateId()
	fmt.Println("ID PARKIR : ", id)
	req := map[string]interface{}{
		"indParking":  id,
	}

	bytesRepresentation, err := json.Marshal(req)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost:8080/api/parkir/in", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.Body)

}

func main() {
	var parkir Parkir
	// scanner
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
		fmt.Println("99. exit")

		fmt.Print("input pilihan :")
		scanner.Scan()
		input = scanner.Text()
		switch input {
		case "1":
			fmt.Println("\nSILAHKAN MASUK JURAGAN...!!!\n")
			parkir.SetMasuk()
			break
		default:
			fmt.Println("Something when wrong..!!")
		}

	}
}
