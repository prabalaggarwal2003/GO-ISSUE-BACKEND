package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){


	d, err := gorm.Open("mysql", "unv1fcsdpjrssiks:G5TzMLd5Y7TmJHnp4k3e@tcp(127.0.0.1:3306)/issue")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}