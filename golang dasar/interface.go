package main

import "fmt"

type hasName interface {
	GetName() string
}

func sayHello(hasName hasName) {
	fmt.Println("Hello", hasName.GetName())
}

func main() {
	sayHello("Fahri")
}
