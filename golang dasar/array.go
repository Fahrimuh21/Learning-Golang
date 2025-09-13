package main

import "fmt"

func main() {
	var names [3]string
	//[3]array string
	names[0] = "fahri"
	names[1] = "fahrimuh"
	names[2] = "muhammad"

	fmt.Println(names[0])

	var values = [...]int{
		//bisa kaya gitu
		90,
		10,
		5,
	}
	fmt.Println(values)
	fmt.Println(len(values))

	var arr2 = [3]int{1, 2, 3}
	arr3 := [3]string{"Go", "Lang", "Dev"}

	fmt.Println(arr2[1])
	fmt.Println(arr3[2])

	//index tertentu
	arr6 := [5]int{0: 100, 3: 200}
	fmt.Println(arr6) // [100 0 0 200 0]

	//array 2 dimensi
	var matrix [2][3]int
	matrix[0][1] = 10

	arr7 := [2][2]string{
		{"a", "b"},
		{"c", "d"},
	}
	fmt.Println(arr7[1][1])

	//akses lewat pointer
	arr := [3]int{10, 20, 30}
	ptr := &arr   // pointer ke array arr

	(*ptr)[0] = 99   // akses elemen array lewat pointer

	fmt.Println("arr:", arr)   // [99 20 30]
	fmt.Println("ptr:", *ptr)  // [99 20 30]


}
