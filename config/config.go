package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	dsn := "user:password@tcp(127.0.0.1:3306)/todo_app"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal membuka koneksi database:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}
	fmt.Println("Database terkoneksi!")
	return db
}
