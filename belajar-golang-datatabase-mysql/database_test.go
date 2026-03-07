package belajar_golang_database

import (
	"database/sql" //package sql adalah package yang digunakan untuk mengakses database SQL seperti MySQL, PostgreSQL, SQLite, dll
	"testing"

	_ "github.com/go-sql-driver/mysql" //import mysql driver kalo ga dipanggil ga bakal bisa konek ke mysql
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database") //akses database mysql dengan user root, password kosong, host localhost, port 3306, dan nama database belajar_golang_database
	if err != nil {
		panic(err) //jika terjadi error maka program akan berhenti dan menampilkan error
	}
	defer db.Close() //untuk memastikan koneksi ke database ditutup setelah selesai digunakan

	// gunakan DB
}
