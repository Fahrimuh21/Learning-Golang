package main

import "fmt"

//nil = empty value
// Nil adalah nilai default untuk beberapa tipe
// data di Go yang menunjukkan bahwa variabel
// tersebut tidak memiliki nilai atau referensi apapun.
// Tipe data yang bisa bernilai nil antara lain:
func main() {
	// Pointer
	var ptr *int
	fmt.Println("Pointer:", ptr) // nil

	// Slice
	var numbers []int
	fmt.Println("Slice:", numbers)                // []
	fmt.Println("Slice == nil ?", numbers == nil) // true

	// Map
	var data map[string]int
	fmt.Println("Map:", data)                // map[]
	fmt.Println("Map == nil ?", data == nil) // true

	// Channel
	var ch chan int
	fmt.Println("Channel:", ch) // nil

	// Function
	var fn func()
	fmt.Println("Function:", fn) // nil

	// Interface
	var anything interface{}
	fmt.Println("Interface:", anything) // <nil>
}
