package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Masuk struct {
	IDPARKIR    string    `json:"indParking"`
	TANGGALMASUK  time.Time `json:"tanggalMasuk"`

}

type Keluar struct {
	TANGGALKELUAR  time.Time `json:"tanggalMasuk"`
	PLAT          string    `json:"plat"`
	TIPE          string    `json:"tipe"`
}

type Parking struct {
	MASUK *Masuk
	KELUAR *Keluar
	TARIF int
}


var ArrayMasuk []Masuk
var ArrayKeluar []Keluar
var parking []Parking

func setParking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var result map[string]interface{}
	json.NewDecoder(r.Body).Decode(&result)
	id := fmt.Sprintf("%v", result["indParking"])
	start := GetTime()
	object := Masuk{id,start}
	ArrayMasuk= append(ArrayMasuk, object)
	log.Println(object)

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



func getParkingId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range ArrayMasuk {
		if item.IDPARKIR == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Masuk{})
}


func outParking(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	//var masuk keluar
	var jsonR = Parking{}
	err = json.Unmarshal(body, &jsonR)
	parking = append(parking, jsonR)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(&Parking{})
}

func getParkingIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ArrayMasuk)
}

func getParkingOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ArrayKeluar)
}

func main() {
	fmt.Print("server is running in port 8080 ")
	router := mux.NewRouter()
	router.HandleFunc("/api/list/in", getParkingIn).Methods("GET")
	router.HandleFunc("/api/list/{id}", getParkingId).Methods("GET")
	router.HandleFunc("/api/list/out", getParkingOut).Methods("GET")
	router.HandleFunc("/api/parkir/in", setParking).Methods("POST")
	router.HandleFunc("/api/parkir/out/{id}", outParking).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))


}
