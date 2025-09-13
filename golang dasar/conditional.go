package main

import "fmt"

func main() {
	name := "fahri"

	// if (name == "fahri") {
	// 	fmt.Println("hello fahri ")
	// } else {
	// 	fmt.Println()
	// }

	//short statetment
	if length := len(name); length > 5 {
		fmt.Println("name terlalu panjang")
	} else {
		fmt.Println("nama sudah bernar")
	}

}
