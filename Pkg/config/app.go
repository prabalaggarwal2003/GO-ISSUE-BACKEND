package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	dbName := os.Getenv("MYSQL_ADDON_DB")
	dbUser := os.Getenv("MYSQL_ADDON_USER")
	dbPass := os.Getenv("MYSQL_ADDON_PASSWORD")
	dbHost := os.Getenv("MYSQL_ADDON_HOST")
	dbPort := os.Getenv("MYSQL_ADDON_PORT")

	d, err := gorm.Open("mysql://"+ dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
