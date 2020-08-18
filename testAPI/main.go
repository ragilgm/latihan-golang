package main

import (
	"encoding/json"
	"fmt"
	"github.com/ragilmaulana/restapi/tugas-golang/testAPI/entities/city"
	provinsi "github.com/ragilmaulana/restapi/tugas-golang/testAPI/entities/provinsi"
	"io/ioutil"
	"net/http"
)

//func main() {
//	////SeachProvince()
//	 result,_ := searchKota("2")
//	 var kota_id string
//	 var provinsi_id string
//for _, value := range result {
//	fmt.Println(value.City_Id,value.City_Name)
//	kota_id = value.City_Id
//	provinsi_id = value.Province_Id
//}
//
//
//	DetailKota(kota_id,provinsi_id)
//
//
//}

type Provinsi struct {
	province_id int `json:"province_id"`
	province    string `json:"province"`
}

func main() {


}

//func main() {
//
//	url := "https://api.rajaongkir.com/starter/cost"
//
//	payload := strings.NewReader("origin=501&destination=114&weight=1700&courier=jne")
//
//	req, _ := http.NewRequest("POST", url, payload)
//
//	req.Header.Add("key", "3cc5610862d3339c470a0ad507b9b75a")
//	req.Header.Add("content-type", "application/x-www-form-urlencoded")
//
//	res, _ := http.DefaultClient.Do(req)
//
//	defer res.Body.Close()
//	body, _ := ioutil.ReadAll(res.Body)
//
//	fmt.Println(res)
//	fmt.Println(string(body))
//
//}

func searchKota(provinsi_id string) ([]city.Result, error) {
	url := "https://api.rajaongkir.com/starter/city?province=" + provinsi_id + ""

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("key", "3cc5610862d3339c470a0ad507b9b75a")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	textByte := []byte(body)

	city := city.City{}

	err := json.Unmarshal(textByte, &city)
	if err != nil {
		panic(err)
	}
	return city.RajaOngkir.Result, err
}

func DetailKota(kota_id, provinsi_id string) ([]city.Result, error) {
	url := "https://api.rajaongkir.com/starter/city?id=" + kota_id + "province=" + provinsi_id + ""

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("key", "3cc5610862d3339c470a0ad507b9b75a")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	textByte := []byte(body)

	city := city.City{}

	err := json.Unmarshal(textByte, &city)
	if err != nil {
		panic(err)
	}
	return city.RajaOngkir.Result, err
}

func SeachProvince() {

	url := "https://api.rajaongkir.com/starter/province"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("key", "3cc5610862d3339c470a0ad507b9b75a")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	textByte := []byte(body)

	listProvinsi := provinsi.Provinsi{}

	err := json.Unmarshal(textByte, &listProvinsi)

	if err != nil {
		panic(err)
		return
	}
	for _, value := range listProvinsi.RajaOngkir.Result {
		fmt.Printf("%v,%v", value.Provinsi_Id, value.Province)
		fmt.Println()
	}
}
