package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamingDecoderJSON(t *testing.T) {
	reader, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}

func TestStreamingEncoderJSON(t *testing.T) {
	writer, _ := os.Create("CustomerOut.JSON")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Muhammad",
		LastName:  "Fahri",
		umur:      19,
	}

	encoder.Encode(customer)
}
