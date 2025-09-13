package main

import "fmt"

//variadic berarti inputan nya slice/map/array
// func sum(nums ...int) int {
// 	total := 0
// 	for _, n := range nums {
// 		total += n
// 	}
// 	return total
// }

// //mutasi dan append
// func mutate(nums ...int) {
// 	if len(nums) > 0 {
// 		nums[0] = 99 // mengubah underlying array
// 	}
// 	nums = append(nums, 100) // append pada header nums lokal
// 	fmt.Println("inside:", nums)
// }
// func main() {
// 	fmt.Println(sum(1, 2, 3)) // 6
// 	fmt.Println(sum())        // 0
// 	s := []int{4, 5, 6}
// 	fmt.Println(sum(s...)) // 15
// 	mutate(s...)           // unpack slice
// 	fmt.Println("after:", s)

// }

// func sumAll(numbers ...int) int {
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }

// func main() {
// 	fmt.Println(sumAll(10, 30, 20)) //bisa langsung cui

// 	//untuk variadic di slice
// 	numbers := []int{10, 10, 10, 10}

// 	fmt.Println(sumAll(numbers...))

// }

//menjadikan fungsi sebagai variable
// func getGopdBye(name string) string {
// 	return "good bye" + name
// }

// func main() {
// 	goodbye := getGopdBye
// 	fmt.Println(goodbye("fahri"))
// }

//define
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Eko", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("fahri", filter)

}
