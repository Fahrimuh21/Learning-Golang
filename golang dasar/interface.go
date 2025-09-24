package main

import "fmt"

// Definisi interface
type HasName interface {
	GetName() string
}

// Struct
type Customer struct {
	Name string
}

// Implementasi method
func (c Customer) GetName() string {
	return c.Name
}

func sayHello(h HasName) {
	fmt.Println("Hello", h.GetName())
}

func main() {
	rully := Customer{Name: "Rully"}
	sayHello(rully) // Hello Rully
}
