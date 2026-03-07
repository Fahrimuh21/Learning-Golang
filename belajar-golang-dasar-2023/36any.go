package main

import "fmt"

func Ups() any /*bisa juga bernama interface{}*/ {
	//return 1
	//return true
	return "Ups"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
