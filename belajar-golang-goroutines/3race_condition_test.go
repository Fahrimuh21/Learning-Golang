package belajar_golang_goroutines

//tentang race condition pada goroutine
//dokumentasi race condition: https://golang.org/doc/articles/r

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			//setiap goroutine menambahkan nilai x sebanyak 100 kali
			for j := 1; j <= 100; j++ {
				x = x + 1
				//race condition terjadi karena banyak goroutine yang mengakses variable x secara bersamaan
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
	//Counter =  97654 (hasilnya tidak selalu sama
}
