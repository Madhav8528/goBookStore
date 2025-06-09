package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {

	Db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/book?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = Db
}

func GetDb() *gorm.DB {
	return db
}
