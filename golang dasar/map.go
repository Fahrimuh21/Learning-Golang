package main

import "fmt"

func main() {
	// person := map[string]string{
	// 	"name":    "fahri",
	// 	"address": "Brebes",
	// }
	// fmt.Println(person)
	// fmt.Println(person["name"])
	// fmt.Println(person["address"])

	// var person map[string]string = map[string]string{}
	// person["name"] = "fahri"
	// person["addres"] = "Brebes"

	book := make(map[string]string)
	book["tittle"] = "buku golang"
	book["author"] = "fahri"
	book["wrong"] = "ups"

	delete(book, "wrong")

	fmt.Println(book)

}
