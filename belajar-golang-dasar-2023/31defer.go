package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() //logging dijalankan setelah runApplication
	fmt.Println("Run application")
}

func main() {
	runApplication()
}
