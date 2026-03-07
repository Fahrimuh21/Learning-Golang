package belajar_golang_context

//context adalah package yang digunakan untuk mengatur batas waktu (timeout), pembatalan (cancellation),
//dan juga untuk mengirim data antar goroutine secara aman dan terstruktur.
//jika satu proses batal maka proses turunannya juga akan batal
//mengirim data request ke proses lain biasanya berupa sinyal
//dokumentasi resmi: https://pkg.go.dev/context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	// background adalah context kosong yang bersifat root atau induk dari context lainnya
	// biasanya digunakan pada main function, test function, dan server request
	fmt.Println(background)

	todo := context.TODO()
	// TODO adalah context kosong yang digunakan ketika kita belum tahu context apa yang akan digunakan
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background() //proses utama/induk

	contextB := context.WithValue(contextA, "b", "B") // membuat context baru dari contextA dengan menambahkan value "b":"B"
	contextC := context.WithValue(contextA, "c", "C") // membuat context baru dari contextA dengan menambahkan value "c":"C"

	contextD := context.WithValue(contextB, "d", "D") //parrent contextB
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f")) //naik ke parrent jika tidak ada di contextF
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextA.Value("b"))
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				//jika ada sinyal cancel dari context
				return
			default:
				destination <- counter //mengirim data ke channel
				counter++
				time.Sleep(1 * time.Second) // simulasi slow
			}
		}
	}()

	return destination
}

// contoh penggunaan context dengan cancel, timeout, dan deadline
// untuk mengatur pembatalan proses goroutine
// dokumentasi resmi: https://pkg.go.dev/context
func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel() // mengirim sinyal cancel ke context

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

// contoh penggunaan context dengan timeout
// untuk mengatur batas waktu proses goroutine
// jika tidak memenuhi waktu yang ditentukan maka proses akan dibatalkan/cancel
// dokumentasi resmi: https://pkg.go.dev/context
func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel() // untuk memastikan cancel dipanggil agar tidak terjadi memory leak

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}


// contoh penggunaan context dengan deadline
//beda dengan timeout, deadline menentukan waktu tertentu
// untuk mengatur batas waktu proses goroutine hingga waktu tertentu
// jika tidak memenuhi waktu yang ditentukan maka proses akan dibatalkan/cancel
// dokumentasi resmi: https://pkg.go.dev/context
func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
