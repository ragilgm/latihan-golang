package entities

type User struct {
	Id_user  int
	Password string
	Nama     string
	Role     string
	Cabang   string
}

type Nasabah struct {
	CIF     int
	NIK     int
	Nama    string
	TTL     string
	TL      string
	Alamat  string
	No_Telp string
}

type SetorTunai struct {
	No_req int
	Nominal int
	Berita string
}


type NasabahDetail struct {
	CIF    int
	No_Req int
	Saldo  int
}

type Transaksi struct {
	id_transaksi    int
	id_user         int
	no_rekening     int
	tanggal         string
	jenis_transaksi string
	nominal         int
	saldo           int
	berita          string
}
