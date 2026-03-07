package belajar_golang_goroutines

//tentang penggunaan sync.Cond pada goroutine
//dokumentasi sync.Cond: https://pkg.go.dev/sync#Cond

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}

// membuat locker untuk sync.Cond
// berguna untuk mengunci akses ke resource bersama
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)
	cond.L.Lock() // mengunci sebelum menunggu kondisi, L adalah locker
	cond.Wait()   // menunggu sinyal dari goroutine lain
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
			//memberi sinyal ke satu goroutine yang menunggu kondisi
			//satu per satu
		}
	}()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// 	//memberi sinyal ke semua goroutine yang menunggu kondisi
	// 	//langsung semua
	// }()

	group.Wait()
}
