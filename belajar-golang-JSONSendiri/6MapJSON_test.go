package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P0001", "name":"Muhammad Fahri", "umur":"19"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["name"])

}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":   "P0001",
		"name": "Muhammad Fahri",
		"umur": "19",
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))

}
