package main

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover() //subscribe dari panic
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error") //publish ke recover
	}
}

func main() {
	runApp(true)
	fmt.Println("Eko Kurniawan Khannedy")
}
