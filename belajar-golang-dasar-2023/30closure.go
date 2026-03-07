package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	} //buat ngubah data counter bisa

	increment()
	increment()
	increment()

	fmt.Println(counter)
}
