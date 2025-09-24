package main

import "fmt"

// Fungsi dengan parameter interface kosong
// Bisa menerima tipe data apapun
func printAnything(value interface{}) {
	fmt.Println("Value:", value)
}

// Fungsi dengan type assertion
func getString(value interface{}) {
	// cek apakah value bertipe string
	str, ok := value.(string)
	if ok {
		fmt.Println("String value:", str)
	} else {
		fmt.Println("Bukan string")
	}
}

// Fungsi dengan type switch
func checkType(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Println("Tipe data string dengan nilai:", v)
	case int:
		fmt.Println("Tipe data int dengan nilai:", v)
	case bool:
		fmt.Println("Tipe data bool dengan nilai:", v)
	default:
		fmt.Println("Tipe data tidak diketahui")
	}
}

func main() {
	// interface kosong bisa menampung apa saja
	var data interface{}

	data = "Golang"
	printAnything(data) // Value: Golang
	getString(data)     // String value: Golang
	checkType(data)     // Tipe data string dengan nilai: Golang

	data = 123
	printAnything(data) // Value: 123
	getString(data)     // Bukan string
	checkType(data)     // Tipe data int dengan nilai: 123

	data = true
	printAnything(data) // Value: true
	checkType(data)     // Tipe data bool dengan nilai: true
}
