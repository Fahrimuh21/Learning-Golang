package main

import "fmt"

func faktorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * faktorial(n-1)
	}
}
func main() {
	fmt.Println(faktorial(5))
}
