package provinsi

type Provinsi struct {
	RajaOngkir Header `json:"rajaongkir"`
}
type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type Header struct {
	Query  []Query  `json:"query"`
	Status Status   `json:"status"`
	Result []Result `json:"results"`
}

type Query struct {
	Id int `json:"id"`
}

type Result struct {
	Provinsi_Id string `json:"province_id"`
	Province    string `json:"province"`
}
