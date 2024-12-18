package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root@tcp(127.0.0.1:3306)/sistem_tracking?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal menghubungkan ke database: %v", err)
	}

	fmt.Println("Berhasil terhubung ke database menggunakan GORM!")
}
