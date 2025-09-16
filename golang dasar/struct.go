package main

import "fmt"

type Customer struct {
	Name, Adress string
	Age          int
}

func main() {
	var fahri Customer
	fahri.Name = "Muhammad Fahri"
	fahri.Adress = "Indonesia"
	fahri.Age = 18

	joko := Customer{
		Name:   "Joko",
		Adress: "Indonesia",
		Age:    18,
	}

	fmt.Println(fahri)
	fmt.Println(joko)

}
