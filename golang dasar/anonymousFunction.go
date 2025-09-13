package main

import "fmt"

type BlackList func(string) bool

func RegisterUser(name string, blaclist BlackList) {
	if blaclist(name) {
		fmt.Println("you are blocked ", name)
	} else {
		fmt.Println("welcome", name)
	}

}

func main() {
	RegisterUser("anjing", func(name string) bool {
		return name == "anjing"
	})//langsung masukin aja
}
