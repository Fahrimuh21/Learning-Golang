package belajar_golang_goroutines

//tentang penggunaan sync.Once pada goroutine
//dokumentasi sync.Once: https://pkg.go.dev/sync#Once

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			//goroutine hanya akan menjalankan fungsi OnlyOnce satu kali saja
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}
