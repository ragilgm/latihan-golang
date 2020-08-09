package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	)

func GetMysqlDB()(db *sql.DB, er error){
	db, err := sql.Open("mysql", "root:Ragil404*@/bds")
	if err != nil {
		panic(err)
	}
	return
}
