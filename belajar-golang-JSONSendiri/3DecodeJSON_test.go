package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// disini tuh berarti untuk membuat objeknya sendiri
// type Customer struct {
// 	FirstName string
// 	LastName  string
// 	umur      int
// }

func TestDecodeJson(t *testing.T) {
	jsonRequest := `{"FirstName":"Muhammad", "LastName":"Fahri"}`
	jsonBytes := []byte(jsonRequest)
	customer := &Customer{}

	//menyimpan decode ke customer dengan customer itu interface
	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Print(customer.FirstName, " ")
	fmt.Println(customer.LastName)

}
