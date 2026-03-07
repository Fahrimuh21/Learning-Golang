package main

import "fmt"

type Filter func(string) string //gunanya agar untuk menamai fungsi sebagai parameter

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Eko", spamFilter) //manggil fungsi spamfilter sebagai filter yaa

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
