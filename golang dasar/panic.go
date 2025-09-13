package main

import "fmt"

func endApp() {
	fmt.Println("end app")
}

func runApp(error bool) {
	if error {
		panic("ups eror")

	}
	endApp()

}

func main() {
	runApp(true)
}
