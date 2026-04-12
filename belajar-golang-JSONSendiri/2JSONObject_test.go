package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

// disini tuh berarti untuk membuat objeknya sendiri
type Customer struct {
	FirstName string
	LastName  string
	umur      int
	Hobbies   []string
	Addresses []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Muhammad",
		LastName:  "Fahri",
		umur:      19,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
