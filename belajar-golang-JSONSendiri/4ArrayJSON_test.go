package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName: "Muhammad",
		LastName:  "Fahri",
		Hobbies:   []string{"Game", "Ngeteh", "Ngoding"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Muhammad","LastName":"Fahri","Hobbies":["Game","Ngeteh","Ngoding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName, " ", customer.LastName)
	fmt.Print(customer.Hobbies)
}

func TestJsonArrayKompleks(t *testing.T) {
	customer := Customer{
		FirstName: "Muhammad",
		Addresses: []Address{
			{
				Street:     "Jalan Belum Ada",
				Country:    "Indonesia",
				PostalCode: "1111",
			},
			{
				Street:     "Jalan lagi dibangun",
				Country:    "Indonesia",
				PostalCode: "9999",
			},
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayKompleksDecode(t *testing.T) {
	jsonString := `{"FirstName":"Muhammad","LastName":"","Hobbies":null,"Addresses":[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"1111"},{"Street":"Jalan lagi dibangun","Country":"Indonesia","PostalCode":"9999"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.Addresses)
}

func TestJsonArrayKompleksDecodeOnlyJsonArray(t *testing.T) {
	jsonString := `[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"1111"},{"Street":"Jalan lagi dibangun","Country":"Indonesia","PostalCode":"9999"}]`
	jsonBytes := []byte(jsonString)

	Addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, Addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(Addresses)
}
