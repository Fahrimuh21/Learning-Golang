package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runGetApp() {
	defer logging() // ini buat skip paling akhir program
	//buat mengatasi eror kalo ada di atasnya
	fmt.Println("run Application")
}
func main() {
	runGetApp()
}
