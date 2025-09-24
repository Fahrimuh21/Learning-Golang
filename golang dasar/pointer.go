package main

import "fmt"

type address struct {
	City, Province, Country string
}

func main() {
	address1 := address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := address1

	address2.City = "Bandung"

	fmt.Println(address1) // address1 ga keubah
	fmt.Println(address2) // address2 keubah
}
