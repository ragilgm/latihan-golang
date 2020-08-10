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
	Tempat_Lahir     string
	Tanggal_Lahir string
	Alamat  string
	No_Telp string
}

type NasabahDetail struct {
	CIF    int
	No_Req int
	Saldo  int
}

type NasabahInfo struct {
	Nasabah Nasabah
	NasabahDetail NasabahDetail
}

type Transaksi struct {
	Id_Transaksi    int
	Id_User         int
	No_Rekening     int
	Tanggal         string
	Jenis_Transaksi string
	Nominal         int
	Saldo           int
	Berita          string
}

type Overbooking struct {
	 IdUser  int
	NASABAH_DETAIL1 NasabahDetail
	NASABAH_DETAIL2 NasabahDetail
	Nominal string
	Berita string
}