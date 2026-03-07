package belajar_golang_goroutines

//tentang penggunaan timer pada goroutine
//dokumentasi time.Timer: https://pkg.go.dev/time#Timer
//dokumentasi time.After: https://pkg.go.dev/time#After
//dokumentasi time.AfterFunc: https://pkg.go.dev/time#AfterFunc

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	//disini +5 detik dari waktu sekarang
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	//disini +5 detik dari waktu sekarang
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())

	group.Wait()
}
