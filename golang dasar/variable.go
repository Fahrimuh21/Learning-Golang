package main

import "fmt"

func main() {
	//versi 1
	var name string
	name = "Muhammad Fahri"
	fmt.Println(name)

	//versi 2
	/* var name = "Muhammad Fahri"
	fmt.Println(name) */

	//versi 3 umum
	//deklarasi awal
	umur := 12
	fmt.Println(umur)

	umur = 18
	fmt.Println(umur)

	//versi 4
	var (
		firstName = "Muhammad"
		lastName  = "Fahri"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
