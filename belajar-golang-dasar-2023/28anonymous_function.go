package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	} //buat ini loh masukin function, kelemahan => dokumentasi kurang rapih
	registerUser("eko", blacklist)

	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	}) //masukin function juga langsung bisa
}
