package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/ragilmaulana/restapi/tugas-golang/BranchDeliverySistem/entities"
	"time"
)

type UserModels struct {
	DB *sql.DB
}

func (um UserModels) Login(nama, password string) (User, error) {
	rows, err := um.DB.Query("SELECT * FROM users WHERE nama LIKE ? AND PASSWORD LIKE ? ", "%"+nama+"%", "%"+password+"%")
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

func (um UserModels) SetorTunai(rekeningTujuan, nominal int, berita string) (NasabahDetail, error) {
	rows, err := um.DB.Query("SELECT * FROM nasabah_detail where no_req=?",rekeningTujuan)
	if err != nil {
		panic(err)
	} else {
		// looping data
		var user NasabahDetail
		for rows.Next() {
			var cif int
			var no_req int
			var saldo int
			err2 := rows.Scan(&cif, &no_req, &saldo)
			if err2 != nil {
				return NasabahDetail{}, err
			} else {
				user = NasabahDetail{
					CIF: cif, No_Req: no_req, Saldo: saldo,
				}
			}
		}

		if user.Saldo <= nominal {
			//	return NasabahDetail{}, err
			//} else {
			//	setor := SetorTunai{
			//		No_req:  user.No_Req,
			//		Nominal: nominal,
			//		Berita:  berita,
			//	}
			//	um.AddSetorTunai(user, setor)
			return NasabahDetail{}, nil
		} else {

			return user, nil
		}
	}
}
func (um UserModels) AddSetorTunai(nasabah NasabahDetail, setor SetorTunai) (int, error) {
	tanggal := time.Now().Format("2006-01-02 15:04:05")
	saldo := nasabah.Saldo + setor.Nominal
	rows, err := um.DB.Exec("insert into transaksi (no_req,tanggal,nominal,saldo,berita) values (?,?,?,?,?)",
		setor.No_req, tanggal, setor.Nominal, saldo, setor.Berita)
	if err != nil {
		return 0, err
	} else {
		idUser, _ := rows.LastInsertId()
		return int(idUser), nil
	}

}

//rows, err := um.DB.Query("SELECT * FROM nasabah_detail")

//err2 := rows.Scan(&cif, &no_req, &saldo)
//if err2 != nil {
//	return entities.NasabahDetail{},err
//} else {
//	fmt.Print(cif)
//	fmt.Print(no_req)
//	fmt.Print(saldo)
//	users := entities.NasabahDetail{CIF: cif,No_Req: no_req,Saldo: saldo}
//	fmt.Println(users)
//	return users,err
//}

//if err2 != nil {
//	return &entities.NasabahDetail{CIF: cif,No_Req: no_req,Saldo: saldo}, err
//} else {
//	if rows == nil {
//		fmt.Println("no req tujuan tidak di temukan")
//		return &entities.NasabahDetail{},err
//	}else if saldo < nominal {
//		fmt.Println("saldo tidak mencukupi")
//		return &entities.NasabahDetail{},err
//	}else {
//		rows, err := um.DB.Exec("insert into nasabah_detail (cif,no_req,saldo) values (?,?,?) ",cif, rekeningTujuan, nominal)
//		if err != nil {
//			return &entities.NasabahDetail{},err
//		}else {
//			row, _ := rows.RowsAffected()
//			fmt.Println("berhasil",row)
//			return &entities.NasabahDetail{
//				CIF: cif,
//				No_Req: rekeningTujuan,
//				Saldo: nominal,
//			},nil
//		}
//
//	}
//}
//}

//
//func (um UserModels) FindALL() ([]entities.User, error) {
//	rows, err := um.DB.Query("select * from user")
//	if err != nil {
//		return nil, err
//	} else {
//		var users []entities.User
//		for rows.Next() {
//			var id int
//			var firstname string
//			var lastname string
//
//			err2 := rows.Scan(&id, &firstname, &lastname)
//			if err2 != nil {
//				return nil, err2
//			} else {
//				user := entities.User{
//					ID: id, FIRSTNAME: firstname, LASTNAME: lastname,
//				}
//				users = append(users, user)
//			}
//		}
//		return users, nil
//	}
//}
//
//

//
//func (um UserModels) Update(user *entities.User) (int64, error) {
//	rows, err := um.DB.Exec("update user set firstName = ?, lastName = ? where id=?", user.FIRSTNAME, user.LASTNAME,user.ID)
//	if err != nil {
//		return 0, err
//	} else {
//		idUser,_ := rows.RowsAffected()
//		return idUser, nil
//	}
//
//}
