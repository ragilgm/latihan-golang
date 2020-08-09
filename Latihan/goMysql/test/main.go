package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Province struct {
	//RAJAONGKIR string
	QUERY []string `json:"query"`
	STATUS *Status `json:"status"`
	RESULT []*Result `json:"results"`
}

type Status struct {
	CODE string `json:"code"`
	DESCRIPTION string `json:"description"`
}


type Result struct {
	provinsi_id string `json:"province_id"`
	provinsi string `json:"provinsi"`
}

//var province []Province

func main() {



	url := "https://api.rajaongkir.com/starter/province"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("key", "3cc5610862d3339c470a0ad507b9b75a")

	res, _ := http.DefaultClient.Do(req)


	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	for _, value := range result {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println()
	}

	//var jsonR = Province{}
	//err := json.Unmarshal(body, &jsonR)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//
	//fmt.Println(jsonR)
	//
	//fmt.Println(res)
	fmt.Println(string(result))

}