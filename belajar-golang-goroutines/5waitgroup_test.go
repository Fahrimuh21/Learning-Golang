package belajar_golang_goroutines

//tentang menunggu goroutine selesai dijalankan
//dokumentasi sync.WaitGroup: https://pkg.go.dev/sync#WaitGroup

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	//ngerunning 1 goroutine

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	//nunggu sampai semua goroutine selesai
	fmt.Println("Selesai")
}
