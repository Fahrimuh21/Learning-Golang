package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = 100 * time.Second     //100 detik
	var duration2 time.Duration = 10 * time.Millisecond //100 milisekon
	var duration3 time.Duration = duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("duration : %d\n", duration3)
}
