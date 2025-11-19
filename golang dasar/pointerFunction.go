package main


import "fmt"

type address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(alamat address){
	address.Country = "Indonesia";
}

func main() {
	address := ChangeAddressToIndonesia("")
}
