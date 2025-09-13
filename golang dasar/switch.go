package main

import "fmt"

func main() {
	name := "fahri"

	switch name {
	case "joko":
		fmt.Println("hello fahri")
	case "fahri":
		fmt.Println("hello joko")
	default:
		fmt.Println("hai boleh kenalan")
	}

}
