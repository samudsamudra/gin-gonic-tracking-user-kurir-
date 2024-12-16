package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	// Ganti dengan konfigurasi database kamu
	dsn := "root@tcp(127.0.0.1:3306)/sistem_tracking"
	var err error

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal terkoneksi ke database: %v", err)
	}

	// Ping database untuk memastikan koneksi berhasil
	if err = DB.Ping(); err != nil {
		log.Fatalf("Tidak bisa melakukan ping ke database: %v", err)
	}

	fmt.Println("Berhasil terhubung ke database!")
}
