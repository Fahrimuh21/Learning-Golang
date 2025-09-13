package main

import "fmt"

func main() {
	// var n int
	// fmt.Scan(n)
	// for i := 0; i <= 100; i++ { //ga boleh ada kurung di for
	// 	fmt.Println(i)
	// }

	num := []int{1, 2, 3}
	for i := range num {
		fmt.Println(i)
	} //index saja 0 1 2

	for i, v := range num {
		fmt.Println("index ", i, "nilai ", v)
	}
	var counter int = 0
	for counter <= 10 {
		counter += 1
	} //kaya while
}
