package helper

import "fmt"

var version = "1.0.0"      //bisa diakses di package ini saja
var Application = "golang" //bisa diakses di package lain

func sayGoodBye(name string) string {
	return "Good Bye " + name
} // private function

func SayHello(name string) string {
	return "Hello " + name
} // public function

func Contoh() {
	sayGoodBye("Eko")
	fmt.Println(version)
} //bisa diakses di package ini saja
