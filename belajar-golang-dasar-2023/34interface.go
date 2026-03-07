package main

import "fmt"

//interface => buat nyimpen fungsi method
//ketika ada dua type struct yang ada mau dimasukin ke dalam satu method maka bisa kita
// gabungkan ke dalam interface
//tapi yang digabungkan itu methodnya
type HasName interface {
	GetName() string
}

//ambil dari interface
func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

//struct 1
type Person struct {
	Name string
}

//method struct 1
func (person Person) GetName() string {
	return person.Name
}

//struct 2
type Animal struct {
	Name string
}

//method struct 2
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Eko"}
	SayHello(person)

	animal := Animal{Name: "Kucing"}
	SayHello(animal)
}
