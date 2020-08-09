package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ragilmaulana/Latihan/goMysql/entities"
)

type UserModels struct {
	DB *sql.DB
}

func (um UserModels) FindALL() ([]entities.User, error) {
	rows, err := um.DB.Query("select * from user")
	if err != nil {
		return nil, err
	} else {
		var users []entities.User
		for rows.Next() {
			var id int
			var firstname string
			var lastname string

			err2 := rows.Scan(&id, &firstname, &lastname)
			if err2 != nil {
				return nil, err2
			} else {
				user := entities.User{
					ID: id, FIRSTNAME: firstname, LASTNAME: lastname,
				}
				users = append(users, user)
			}
		}
		return users, nil
	}
}

func (um UserModels) FIndById(id int) (entities.User, error) {
	rows, err := um.DB.Query("select * from user where id=?", id)
	if err != nil {
		return entities.User{}, err
	} else {
		var user entities.User
		for rows.Next() {
			var id int
			var firstname string
			var lastname string

			err2 := rows.Scan(&id, &firstname, &lastname)
			if err2 != nil {
				return entities.User{}, err
			} else {
				user = entities.User{
					ID: id, FIRSTNAME: firstname, LASTNAME: lastname,
				}
			}
		}
		return user, nil
	}
}

func (um UserModels) Insert(user *entities.User) (int64, error) {
	rows, err := um.DB.Exec("insert into user (firstName,lastName) values (?,?)", user.FIRSTNAME, user.LASTNAME)
	if err != nil {
		return 0, err
	} else {
		idUser,_ := rows.LastInsertId()
		return idUser, nil
	}

}

func (um UserModels) Update(user *entities.User) (int64, error) {
	rows, err := um.DB.Exec("update user set firstName = ?, lastName = ? where id=?", user.FIRSTNAME, user.LASTNAME,user.ID)
	if err != nil {
		return 0, err
	} else {
		idUser,_ := rows.RowsAffected()
		return idUser, nil
	}

}
