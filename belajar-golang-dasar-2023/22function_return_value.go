package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

//multiple values return
func getFullName() (string, string) {
	return "Eko", "Kurniawan"
}

func main() {
	result := getHello("Eko")
	result1, result2 := getFullName()
	firstName, _ := getFullName() //jika hanya butuh 1 nilai return saja
	fmt.Println(firstName)
	fmt.Println(result)
	fmt.Println(result1, result2)

	fmt.Println(getHello("Budi"))
	fmt.Println(getHello("Joko"))
}
