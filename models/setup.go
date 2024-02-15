package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
//sesuaikan dengan kebutuhan
func ConnectDatabase() { //urutan koneksi root:password@tcp(host dan port)/namadb
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/restAPI"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Studi{})

	DB = database
}
