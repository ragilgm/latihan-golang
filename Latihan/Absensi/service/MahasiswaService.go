package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	MahasiswaProto "github.com/ragilmaulana/restapi/tugas-golang/Latihan/Absensi/proto"
)

type Mahasiswa interface {
	FindAll()
	FindById()
	InsertData()
	UpdateData()
}

type MahasiswaService struct {
	DB *sql.DB
}

func (mhs MahasiswaService) FindAll()([]MahasiswaProto.User){
rows, err := mhs.DB.Query("select * from mahasiswa")
if err !=nil {
	panic(err)
}

var users MahasiswaProto.User
for rows.Next(){
var id int64
var F
}
}


}
