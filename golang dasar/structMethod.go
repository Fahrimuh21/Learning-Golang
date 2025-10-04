package main

import "fmt"

type Customer struct {
	Name, Adress string
	Age          int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, My Name is", customer.Name)
}
func main() {
	rully := Customer{Name: "Rully"}
	rully.sayHello()
}
// Output: Hello, My Name is Rully