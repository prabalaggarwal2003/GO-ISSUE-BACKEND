package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){


	d, err := gorm.Open("mysql", "unv1fcsdpjrssiks:G5TzMLd5Y7TmJHnp4k3e@biqvusjvq0ih4w4ikcw3-mysql.services.clever-cloud.com:3306/biqvusjvq0ih4w4ikcw3?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}