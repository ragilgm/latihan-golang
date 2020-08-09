package main

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ragilmaulana/Latihan/goMysql/config"
	"github.com/ragilmaulana/Latihan/goMysql/entities"
	"github.com/ragilmaulana/Latihan/goMysql/models"
)

func main() {
//findAll()
//findById(1)
Insert("sempak","item")

}


func findAll(){
	db, err := config.GetMysqlDB()
	if err != nil {
		panic(err)
	}else {
		user := models.UserModels{
			db,
		}
		users , err := user.FindALL()
		if err != nil {
			panic(err)
		}else {
			for _, value := range users {
				fmt.Printf("%v,%v,%v", value.ID,value.FIRSTNAME, value.LASTNAME)
			}
		}
	}

}

func findById(id int){
	db, err := config.GetMysqlDB()
	if err != nil {
		panic(err)
	}else {
		user := models.UserModels{
			db,
		}
		users , err := user.FIndById(id)
		if err != nil {
			panic(err)
		}else {
				fmt.Printf("%v,%v,%v", users.ID,users.FIRSTNAME, users.LASTNAME)

		}
	}

}

func Insert(firstname, lastname string){
	db, err := config.GetMysqlDB()
	if err != nil {
		panic(err)
	}else {
		models := models.UserModels{
			db,
		}
		user := entities.User{
			FIRSTNAME: firstname,
			LASTNAME: lastname,
		}
		users , err := models.Insert(&user)
		if err != nil {
			panic(err)
		}else {
			fmt.Println("Id",users, "have been inserted")

		}
	}

}