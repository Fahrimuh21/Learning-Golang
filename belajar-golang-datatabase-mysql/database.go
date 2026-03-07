package belajar_golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                  //mengatur minimal koneksi yang terbuka di connection pool
	db.SetMaxOpenConns(100)                 //koneksi maksimal yang boleh terbuka di connection pool
	db.SetConnMaxIdleTime(5 * time.Minute)  //mengatur waktu maksimal koneksi boleh idle sebelum ditutup, jika koneksi idle lebih dari 5 menit maka koneksi akan ditutup (restart koneksi)
	db.SetConnMaxLifetime(60 * time.Minute) //mengatur waktu maksimal koneksi boleh digunakan sebelum ditutup, jika koneksi sudah digunakan lebih dari 60 menit maka koneksi akan ditutup (shutdown koneksi)

	return db
}
