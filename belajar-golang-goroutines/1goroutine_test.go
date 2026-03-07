package belajar_golang_goroutines

//tentang membuat dan menjalankan goroutine
//dokumentasi goroutine: https://golang.org/doc/effective_go#goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // menjalankan goroutine
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
	// memberi waktu agar goroutine selesai sebelum program utama berakhir
	// modul time dokumentasi: https://pkg.go.dev/time
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i) // menjalankan banyak goroutine secara bersamaan
	}

	time.Sleep(5 * time.Second) // memberi waktu agar semua goroutine selesai sebelum program utama berakhir
}
