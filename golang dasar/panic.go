package main

import "fmt"

func endApp() {
	fmt.Println("end app")
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func runApp(error bool) {
	defer endApp() //biar masuk endapp
	if error {
		panic("ups eror") //langsung berhenti
	}

	// endApp() ini tidak akan dipanggil

}

func main() {
	runApp(true)
	fmt.Println("Muhammad Fahri")
}
