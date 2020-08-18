package city

type City struct {
	RajaOngkir RajaOngkir `json:"rajaongkir"`
}

type RajaOngkir struct {
	Query Query `json:"query"`
	Status Status `json:"status"`
	Result []Result `json:"results"`
}

type Query struct {
 Id string `json:"id"`
 Province string `json:"province"`
}

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type Result struct {
	City_Id string `json:"city_id"`
	Province_Id    string `json:"province_id"`
	Province string `json:"province"`
	Type string `json:"type"`
	City_Name string `json:"city_name"`
	Postal_Code string `json:"postal_code"`
}
