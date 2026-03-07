package belajar_golang_goroutines

//kejadian berulang
//dokumentasi time.Ticker: https://pkg.go.dev/time#Ticker
//dokumentasi time.Tick: https://pkg.go.dev/time#Tick

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		//menghentikan ticker setelah 5 detik
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}
	//tiap detik akan mengirim data ke channel C
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)
	//menghasilkan channel yang mengirim data tiap detik

	for time := range channel {
		fmt.Println(time)
	}
}

//tiap detik akan mengirim data ke channel
