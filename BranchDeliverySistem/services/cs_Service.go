package services
//
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/ragilmaulana/restapi/tugas-golang/BranchDeliverySistem/entities"
)

type CustomerServiceModels struct {
	DB *sql.DB
}
// find cif service ==========================================================================================
func (um CustomerServiceModels) FindCif(cif int) (Nasabah, error) {
	rows, err := um.DB.Query("SELECT * FROM nasabah WHERE cif=?", 6666)
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

			err2 := rows.Scan(&cif, &nik, &nama, &tempat_lahir,&tanggal_lahir, &alamat, &no_telp)
			if err2 != nil {
				panic(err2)
			} else {
				nasabah = Nasabah{
					CIF: int(cif),
					NIK: int(nik),
					Nama : nama,
					Tempat_Lahir : tempat_lahir,
					Tanggal_Lahir: tanggal_lahir,
					Alamat: alamat,
					No_Telp: no_telp,
				}
			}
		}
		return nasabah, nil
	}
}
// end find cif service ======================================================================================



// Create Cif ================================================================================================
func (um UserModels) BuatCIF(nasabah Nasabah) (int64, error) {
	rows, err := um.DB.Exec("insert into user (cif,nik,nama,tempat_lahir,tanggal_lahir,alamat,no_telp) values (?,?,?,?,?,?,?)",
		nasabah.CIF, nasabah.NIK,nasabah.Nama,nasabah.Tempat_Lahir,nasabah.Tanggal_Lahir,nasabah.Alamat,nasabah.No_Telp)
	if err != nil {
		return 0, err
	} else {
		idUser,_ := rows.LastInsertId()
		return idUser, nil
	}
}
// ==========================================================================================================

// Create Tabungan ================================================================================================
//func (um UserModels) BuatRekening(cif int, saldoAwal int) (, error) {
//	rows, err := um.DB.Exec("insert into user (cif,nik,nama,tempat_lahir,tanggal_lahir,alamat,no_telp) values (?,?,?,?,?,?,?)",
//		nasabah.CIF, nasabah.NIK,nasabah.Nama,nasabah.Tempat_Lahir,nasabah.Tanggal_Lahir,nasabah.Alamat,nasabah.No_Telp)
//	if err != nil {
//		return 0, err
//	} else {
//		idUser,_ := rows.LastInsertId()
//		return idUser, nil
//	}
//}
// ==========================================================================================================

