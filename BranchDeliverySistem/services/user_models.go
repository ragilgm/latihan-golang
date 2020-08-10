package services

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	BranchDeliverySystem "github.com/ragilmaulana/restapi/tugas-golang/BranchDeliverySistem/bdsProto"
	. "github.com/ragilmaulana/restapi/tugas-golang/BranchDeliverySistem/entities"
	"time"
)

type UserModels struct {
	DB *sql.DB
}

// login =============================================================================================
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

// login =============================================================================================

// check nasabah exist or not ========================================================================
func (um UserModels) FindNoRek(rekeningTujuan int) (NasabahDetail, error) {
	rows, err := um.DB.Query("SELECT * FROM nasabah_detail where no_rekening=?", rekeningTujuan)
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
		"insert into transaksi (id_user, no_rekening,tanggal,jenis_transaksi,nominal,saldo,berita)values(?,?,?,?,?,?,?)",
		userId, nasabah.No_Req, tanggal, jenisTransaksi, nominal, currentSaldo, berita)
	_, err = um.DB.Exec(
		"update nasabah_detail set saldo = ? where no_rekening=?", currentSaldo, nasabah.No_Req)
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
			"insert into transaksi (id_user, no_rekening,tanggal,jenis_transaksi,nominal,saldo,berita)value(?,?,?,?,?,?,?)",
			userId, nasabah.No_Req, tanggal, jenisTransaksi, nominal, currentSaldo, berita)
		_, err = um.DB.Exec(
			"update nasabah_detail set saldo = ? where ne_rekening=?", currentSaldo, nasabah.No_Req)
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
			"insert into transaksi (id_user, no_rekening,tanggal,jenis_transaksi,nominal,saldo,berita)values(?,?,?,?,?,?,?)",
			idUser, rekeningAwal.No_Req, tanggal, jenisTransaksi, nominal, saldoRekAwal, berita)
		if err != nil {
			panic(err)
		}
		fmt.Println(insertRekUtama)
		update, err := um.DB.Exec("update nasabah_detail set saldo = ? where no_rekening=?", saldoRekAwal, rekeningAwal.No_Req)
		if err != nil {
			panic(err)
		}
		fmt.Println(update)
		insertRekKedua, err := um.DB.Exec("insert into transaksi (id_user, no_rekening,tanggal,jenis_transaksi,nominal,saldo,berita)values(?,?,?,?,?,?,?)",
			idUser, rekeingTujuan.No_Req, tanggal, jenisTransaksi, nominal, saldoRekTujuan, berita)
		if err != nil {
			panic(err)
		}
		fmt.Println(insertRekKedua)
		updateRekTujuan, err := um.DB.Exec("update nasabah_detail set saldo = ? where no_rekening=?", saldoRekTujuan, rekeingTujuan.No_Req)
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
	rows, err := um.DB.Query("SELECT * FROM transaksi WHERE no_rekening=?", no_rekening)
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
func (um UserModels) PrintNasabahInfoByRekening(nomor_rekening int) (NasabahInfo, error) {
	rows, err := um.DB.Query("SELECT * FROM nasabah INNER JOIN nasabah_detail ON nasabah.cif = nasabah_detail.cif WHERE no_rekening=?", nomor_rekening)
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
			err2 := rows.Scan(&cif, &nik, &nama, &tempat_lahir, &tanggal_lahir, &alamat, &no_telp, &cif, &no_req, &saldo)
			if err2 != nil {
				panic(err2)
			} else {
				nasabahInfo = NasabahInfo{
					Nasabah: Nasabah{
						CIF:          cif,
						NIK:          nik,
						Nama:         nama,
						Tempat_Lahir: tempat_lahir,
						Alamat:       alamat,
						No_Telp:      no_telp,
					},
					NasabahDetail: NasabahDetail{
						CIF:    cif,
						No_Req: no_req,
						Saldo:  saldo,
					},
				}

			}

		}
		return nasabahInfo, nil
	}
}

//====================================================================================================================

// find nasabah detail by cif ========================================================================================
func (um UserModels) PrintNasabahInfoByCif(cif int) (NasabahInfo, error) {
	rows, err := um.DB.Query("SELECT * FROM nasabah INNER JOIN nasabah_detail ON nasabah.cif = nasabah_detail.cif WHERE nasabah.cif=?", cif)
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
			err2 := rows.Scan(&cif, &nik, &nama, &tempat_lahir, &tanggal_lahir, &alamat, &no_telp, &cif, &no_req, &saldo)
			if err2 != nil {
				panic(err2)
			} else {
				nasabahInfo = NasabahInfo{
					Nasabah: Nasabah{
						CIF:          cif,
						NIK:          nik,
						Nama:         nama,
						Tempat_Lahir: tempat_lahir,
						Alamat:       alamat,
						No_Telp:      no_telp,
					},
					NasabahDetail: NasabahDetail{
						CIF:    cif,
						No_Req: no_req,
						Saldo:  saldo,
					},
				}

			}

		}
		return nasabahInfo, nil
	}
}

//=======================================================================================================================

// find cif service ==========================================================================================
func (um UserModels) FindCIFOrNIK(seach int) (Nasabah, error) {
	rows, err := um.DB.Query("SELECT * FROM nasabah WHERE cif=? OR nik=?", seach, seach)

	if err != nil {
		panic(err)
	} else {
		var nasabah Nasabah
		for rows.Next() {
			var cif int
			var nik int
			var nama string
			var tempat_lahir string
			var tanggal_lahir string
			var alamat string
			var no_telp string

			err2 := rows.Scan(&cif, &nik, &nama, &tempat_lahir, &tanggal_lahir, &alamat, &no_telp)
			if err2 != nil {
				panic(err2)
			}
			nasabah = Nasabah{
				CIF:           int(cif),
				NIK:           int(nik),
				Nama:          nama,
				Tempat_Lahir:  tempat_lahir,
				Tanggal_Lahir: tanggal_lahir,
				Alamat:        alamat,
				No_Telp:       no_telp,
			}
			fmt.Println("called")
			fmt.Println(nasabah)
		}
		return nasabah, nil
	}
}

// check nasabah exist or not ========================================================================
func (um UserModels) FindLastInsertNoRek() (int, error) {
	rows, err := um.DB.Query("SELECT  no_rekening FROM nasabah_detail ORDER BY  no_rekening DESC LIMIT 1")
	if err != nil {
		panic(err)
	} else {
		// looping data
		var no_rekening int
		for rows.Next() {

			err2 := rows.Scan(&no_rekening)
			if err2 != nil {
				panic(err2)
			}
		}
		return no_rekening, nil
	}
}

// end find cif service ======================================================================================

// Create Cif ================================================================================================
func (um UserModels) BuatCIF(nasabah Nasabah) (Nasabah, error) {
	rows, err := um.DB.Exec("insert into nasabah (nik,nama,tempat_lahir,tanggal_lahir,alamat,no_telp) values (?,?,?,?,?,?)",
		nasabah.NIK, nasabah.Nama, nasabah.Tempat_Lahir, nasabah.Tanggal_Lahir, nasabah.Alamat, nasabah.No_Telp)
	fmt.Println("nik", nasabah.NIK)
	if err != nil {
		panic(err)
	} else {
		status, _ := rows.RowsAffected()
		if status > 0 {
			result, err := um.FindCIFOrNIK(nasabah.NIK)
			return result, err
		} else {
			return Nasabah{}, err
		}
	}
}

// ==========================================================================================================

// buat rekening tabungan ===================================================================================
func (um UserModels) BuatTabungan(cif, saldo int64) (NasabahInfo, error) {
	last_no_rekening, _ := um.FindLastInsertNoRek()
	last_no_rekening += 1

	fmt.Println(last_no_rekening)
	rows, err := um.DB.Exec("insert into nasabah_detail (cif,no_rekening,saldo) values (?,?,?)",
		cif, last_no_rekening, saldo)
	if err != nil {
		panic(err)
	}
	status, _ := rows.RowsAffected()
	if status > 0 {
		// call method print nasabah by cif
		result, _ := um.PrintNasabahInfoByCif(int(cif))
		return result, nil
	} else {
		return NasabahInfo{}, nil
	}
}

// =======================================================================================================

// management crud nasabah ====================================================================================
func (um UserModels) UpdateNasabah(nasabah Nasabah) (*BranchDeliverySystem.NASABAH_INFO, error) {
	rows, err := um.DB.Exec("update nasabah set nik = ?, nama = ?, tempat_lahir = ?, tanggal_lahir=?,alamat=?,no_telepon=? where cif =? or nik=?",
		nasabah.NIK, nasabah.Nama, nasabah.Tempat_Lahir, nasabah.Tanggal_Lahir, nasabah.Alamat, nasabah.No_Telp, nasabah.CIF, nasabah.NIK)
	fmt.Println(nasabah)
	if err != nil {
		panic(err)
	}
	fmt.Println("called")
var Respons *BranchDeliverySystem.NASABAH_INFO
	idUser, _ := rows.RowsAffected()
	fmt.Println(idUser)
	if idUser > 0 {

		result, err := um.PrintNasabahInfoByCif(nasabah.CIF)
		if err != nil {
			panic(err)
		}
		Respons =  &BranchDeliverySystem.NASABAH_INFO{
			NASABAH: &BranchDeliverySystem.NASABAH{
				CIF: int64(result.Nasabah.CIF),
				NIK: int64(result.Nasabah.NIK),
				NAMA: result.Nasabah.Nama,
				ALAMAT: result.Nasabah.Alamat,
				TEMPAT_LAHIR: result.Nasabah.Tempat_Lahir,
				TANGGAL_LAHIR: result.Nasabah.Tanggal_Lahir,
				NO_TELP: result.Nasabah.No_Telp,
			},
			NASABAH_DETAIL: &BranchDeliverySystem.NASABAH_DETAIL{
				CIF: int64(result.NasabahDetail.CIF),
				NO_REKENING: int64(result.NasabahDetail.No_Req),
				SALDO: int64(result.NasabahDetail.Saldo),
			},
		}
	}
	return Respons,nil
}
// ============================================================================================================
