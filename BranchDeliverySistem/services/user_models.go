package services

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/ragilmaulana/restapi/tugas-golang/BranchDeliverySistem/entities"
	"time"
)

type UserModels struct {
	DB *sql.DB
}

func (um UserModels) Login(idUser int, password string) (User, error) {
	rows, err := um.DB.Query("SELECT * FROM users WHERE id_user=? AND PASSWORD LIKE ? ", idUser, "%"+password+"%")
	if err != nil {
		return User{}, err
	} else {
		var user User
		for rows.Next() {
			var id_user int
			var password string
			var nama string
			var role string
			var cabang string
			err2 := rows.Scan(&id_user, &password, &nama, &role, &cabang)
			if err2 != nil {
				return User{}, err
			} else {
				user = User{
					Id_user: id_user, Password: password, Nama: nama, Role: role, Cabang: cabang,
				}
			}
		}
		return user, nil
	}
}

// check nasabah exist or not ========================================================================
func (um UserModels) FindNoRek(rekeningTujuan int) (NasabahDetail, error) {
	rows, err := um.DB.Query("SELECT * FROM nasabah_detail where no_req=?", rekeningTujuan)
	if err != nil {
		panic(err)
	} else {
		// looping data
		var nasabahDetail NasabahDetail
		for rows.Next() {
			var cif int
			var no_req int
			var saldo int
			err2 := rows.Scan(&cif, &no_req, &saldo)
			if err2 != nil {
				return NasabahDetail{}, err
			} else {
				nasabahDetail = NasabahDetail{
					CIF: cif, No_Req: no_req, Saldo: saldo,
				}

			}

		}
		return nasabahDetail, nil
	}
}
//===================================================================================================

// service setor tunai to db ==================================================================================
func (um UserModels) SetorTunaiService(userId int, nasabah NasabahDetail, berita string, nominal int) (int, error) {
	tanggal := time.Now().Format("2006-01-02 15:04:05")
	saldo := nasabah.Saldo
	jenisTransaksi := "st"
	currentSaldo := saldo + nominal
	rows, err := um.DB.Exec(
		"insert into transaksi (id_user, no_req,tanggal,jenis_transaksi,nominal,saldo,berita)values(?,?,?,?,?,?,?)",
		userId, nasabah.No_Req, tanggal, jenisTransaksi, nominal, currentSaldo, berita)
	_, err = um.DB.Exec(
		"update nasabah_detail set saldo = ? where no_req=?", currentSaldo, nasabah.No_Req)
	if err != nil {
		return 0, err
	} else {
		idUser, _ := rows.RowsAffected()
		return int(idUser), nil
	}
}
// end of service store tunai ================================================================================

// insert setor tunai to db ==================================================================================
func (um UserModels) TarikTunaiService(userId int, nasabah NasabahDetail, berita string, nominal int) (int, error) {
	tanggal := time.Now().Format("2006-01-02 15:04:05")
	saldo := nasabah.Saldo
	jenisTransaksi := "tt"
	// check saldo overload
	if saldo < nominal {
		return 0, nil
	} else {
		currentSaldo := nasabah.Saldo - nominal
		rows, err := um.DB.Exec(
			"insert into transaksi (id_user, no_req,tanggal,jenis_transaksi,nominal,saldo,berita)value(?,?,?,?,?,?,?)",
			userId, nasabah.No_Req, tanggal, jenisTransaksi, nominal, currentSaldo, berita)
		_, err = um.DB.Exec(
			"update nasabah_detail set saldo = ? where no_req=?", currentSaldo, nasabah.No_Req)
		if err != nil {
			return 0, err
		} else {
			idUser, _ := rows.RowsAffected()
			return int(idUser), nil
		}
	}
}
// end of insert service =====================================================================================

// service overbooking =======================================================================================
func (um UserModels) Overbooking(idUser int, rekeningAwal, rekeingTujuan NasabahDetail, nominal int, berita string) (int, error) {
	tanggal := time.Now().Format("2006-01-02 15:04:05")
	// check saldo overload
	if rekeningAwal.Saldo < nominal {
		return 0, nil
	} else {
		saldoRekAwal := int(rekeningAwal.Saldo - nominal)
		fmt.Println(saldoRekAwal)
		saldoRekTujuan := int(rekeingTujuan.Saldo + nominal)
		fmt.Println(saldoRekTujuan)
		jenisTransaksi := "pb"
		insertRekUtama, err := um.DB.Exec(
			"insert into transaksi (id_user, no_req,tanggal,jenis_transaksi,nominal,saldo,berita)values(?,?,?,?,?,?,?)",
			idUser, rekeningAwal.No_Req, tanggal, jenisTransaksi, nominal, saldoRekAwal, berita)
		if err != nil {
			panic(err)
		}
		fmt.Println(insertRekUtama)
		update, err := um.DB.Exec("update nasabah_detail set saldo = ? where no_req=?", saldoRekAwal, rekeningAwal.No_Req)
		if err != nil {
			panic(err)
		}
		fmt.Println(update)
		insertRekKedua, err := um.DB.Exec("insert into transaksi (id_user, no_req,tanggal,jenis_transaksi,nominal,saldo,berita)values(?,?,?,?,?,?,?)",
			idUser, rekeingTujuan.No_Req, tanggal, jenisTransaksi, nominal, saldoRekTujuan, berita)
		if err != nil {
			panic(err)
		}
		fmt.Println(insertRekKedua)
		updateRekTujuan, err := um.DB.Exec("update nasabah_detail set saldo = ? where no_req=?", saldoRekTujuan, rekeingTujuan.No_Req)
		if err != nil {
			panic(err)
		}
		fmt.Println(updateRekTujuan)

	}
	return idUser, nil
}
// end service overbooking ===================================================================================

// service cetak buku ========================================================================================
func (um UserModels) CetakBuku(no_rekening int) ([]Transaksi, error) {
	rows, err := um.DB.Query("SELECT * FROM transaksi WHERE no_req=?", no_rekening)
	if err != nil {
		return []Transaksi{}, err
	} else {
		var transaksi []Transaksi
		for rows.Next() {
			var id_transaksi int
			var id_user int
			var no_rek int
			var tanggal string
			var jenis_transaksi string
			var nominal int
			var saldo int
			var berita string
			err2 := rows.Scan(&id_transaksi, &id_user, &no_rek, &tanggal, &jenis_transaksi, &nominal, &saldo, &berita)
			if err2 != nil {
				return []Transaksi{}, err
			} else {
				trx := Transaksi{

					Id_Transaksi: id_transaksi,
					Id_User:      id_user, No_Rekening: no_rek,
					Tanggal:         tanggal,
					Jenis_Transaksi: jenis_transaksi,
					Nominal:         nominal,
					Saldo:           saldo,
					Berita:          berita,
				}
				transaksi = append(transaksi, trx)
			}
		}
		return transaksi, nil
	}
}
// end service cetak buku ====================================================================================


// find nasabah by no req ====================================================================================
func (um UserModels) PrintNasabahInfoByRekening(nomor_rekening int) (NasabahInfo, error){
	rows, err := um.DB.Query("SELECT * FROM nasabah INNER JOIN nasabah_detail ON nasabah.cif = nasabah_detail.cif WHERE no_req=?", nomor_rekening)
	if err != nil {
		panic(err)
	} else {
		// looping data
		var nasabahInfo NasabahInfo
		for rows.Next() {
			var cif int
			var nik int
			var nama string
			var tempat_lahir string
			var tanggal_lahir string
			var alamat string
			var no_telp string
			var no_req int
			var saldo int
			err2 := rows.Scan(&cif, &nik, &nama, &tempat_lahir, &tanggal_lahir, &alamat, &no_telp,&cif,&no_req, &saldo)
			if err2 != nil {
				panic(err2)
			} else {
				nasabahInfo = NasabahInfo{
					Nasabah: Nasabah{
						CIF: cif,
						NIK: nik,
						Nama: nama,
						Tempat_Lahir: tempat_lahir,
						Alamat: alamat,
						No_Telp: no_telp,
					},
					NasabahDetail: NasabahDetail{
						CIF: cif,
						No_Req: no_req,
						Saldo: saldo,
					},
				}

			}

		}
		return nasabahInfo, nil
	}
}
//====================================================================================================================


// find nasabah detail by cif ========================================================================================
func (um UserModels) PrintNasabahInfoByCif(cif int) (NasabahInfo, error){
	rows, err := um.DB.Query("SELECT * FROM nasabah INNER JOIN nasabah_detail ON nasabah.cif = nasabah_detail.cif WHERE cif=?", cif)
	if err != nil {
		panic(err)
	} else {
		// looping data
		var nasabahInfo NasabahInfo
		for rows.Next() {
			var cif int
			var nik int
			var nama string
			var tempat_lahir string
			var tanggal_lahir string
			var alamat string
			var no_telp string
			var no_req int
			var saldo int
			err2 := rows.Scan(&cif, &nik, &nama, &tempat_lahir, &tanggal_lahir, &alamat, &no_telp,&cif,&no_req, &saldo)
			if err2 != nil {
				panic(err2)
			} else {
				nasabahInfo = NasabahInfo{
					Nasabah: Nasabah{
						CIF: cif,
						NIK: nik,
						Nama: nama,
						Tempat_Lahir: tempat_lahir,
						Alamat: alamat,
						No_Telp: no_telp,
					},
					NasabahDetail: NasabahDetail{
						CIF: cif,
						No_Req: no_req,
						Saldo: saldo,
					},
				}

			}

		}
		return nasabahInfo, nil
	}
}
//=======================================================================================================================
