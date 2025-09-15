package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runGetApp() {
	defer logging() // ini buat skip paling akhir program
	fmt.Println("run Application")
}
func main() {
	runGetApp()
}
