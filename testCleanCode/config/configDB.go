package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/ragilmaulana/Latihan/clean/pkg/entity"
)

type DbConfig struct {
	 DB *gorm.DB
}
var err error
func ( db *DbConfig)Connect() (*gorm.DB, error) {
	db.DB, err = gorm.Open("mysql", "root:Ragil404*@tcp(127.0.0.1:3306)/sample?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	db.DB.AutoMigrate(&entity.Person{})
	return db.DB, nil
}