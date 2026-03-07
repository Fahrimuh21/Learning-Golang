package main

import (
	"belajar-golang-dasar/database"   //import package database untuk mengakses init di mysql.go dan function GetDatabase
	_ "belajar-golang-dasar/internal" //import untuk menjalankan init di package internal tidak perlu mengakses isinya
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
