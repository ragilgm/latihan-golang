package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	UserService "github.com/ragilmaulana/restapi/tugas-golang/Latihan/goMysql/protoUser"
)

type UserModels struct {
	DB *sql.DB
}

func (um UserModels) FindALL() ([]UserService.User, error) {
	rows, err := um.DB.Query("select * from user")
	if err != nil {
		panic(err)
	} else {
		var users []UserService.User
		for rows.Next() {
			var id int
			var firstname string
			var lastname string

			err2 := rows.Scan(&id, &firstname, &lastname)
			if err2 != nil {
				panic(err2)
			} else {
				user := UserService.User{
					ID: int64(id),
					FIRSTNAME: firstname,
					LASTNAME: lastname,
				}
				users = append(users, user)
			}
		}
		return users, nil
	}
}

func (um UserModels) FIndById(id int) (UserService.User, error) {
	rows, err := um.DB.Query("select * from user where id=?", id)
	if err != nil {
		panic(err)
	} else {
		var user UserService.User
		for rows.Next() {
			var id int
			var firstname string
			var lastname string

			err2 := rows.Scan(&id, &firstname, &lastname)
			if err2 != nil {
				panic(err)
			} else {
				user = UserService.User{
					ID: int64(id),
					FIRSTNAME: firstname,
					LASTNAME: lastname,
				}
			}
		}
		return user, nil
	}
}

func (um UserModels) Insert(user *UserService.User) (int64, error) {
	rows, err := um.DB.Exec("insert into user (firstName,lastName) values (?,?)", user.FIRSTNAME, user.LASTNAME)
	if err != nil {
		return 0, err
	} else {
		idUser,_ := rows.LastInsertId()
		return idUser, nil
	}

}

func (um UserModels) Update(user *UserService.User) (int64, error) {
	rows, err := um.DB.Exec("update user set firstName = ?, lastName = ? where id=?", user.FIRSTNAME, user.LASTNAME,user.ID)
	if err != nil {
		return 0, err
	} else {
		idUser,_ := rows.RowsAffected()
		return idUser, nil
	}

}
