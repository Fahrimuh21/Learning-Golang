package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runGetApp() {
	defer logging()
	fmt.Println("run Application")
}
func main() {
	runGetApp()
}
