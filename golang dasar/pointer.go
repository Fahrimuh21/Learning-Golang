package main

import "fmt"

type address struct {
	City, Province, Country string
}

func main() {
	var address1 address = address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address2 *address = &address1 //copy

	address2.City = "Bandung"
	fmt.Println(address1) // address1 keubah
	fmt.Println(address2)
	*address2 = address{"Jakarta", "Bandung", "Brebes"} // address1 ikut keubah
	fmt.Println(address2)
	address2 = &address{"Jakarta", "Bandung", "Tegal"} //agar tidak keubah address1
	fmt.Println(address1)
	fmt.Println(address2)

	//tidak di deklarasi awal
	alamat1 := new(address)
	alamat2 := alamat1 //otomatis ke pointer
	alamat2.Country = "Indonesia"
	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
