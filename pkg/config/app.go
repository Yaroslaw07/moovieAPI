package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var (
	db *gorm.DB
)

func Connect() {

	//add path
	d, err := gorm.Open("mysql", "root:0707@tcp(127.0.0.1:3306)/moovies?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
