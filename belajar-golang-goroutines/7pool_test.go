package belajar_golang_goroutines

//tentang penggunaan sync.Pool pada goroutine
//dokumentasi sync.Pool: https://pkg.go.dev/sync#Pool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
		//fungsi yang akan dijalankan ketika pool kosong
	}

	pool.Put("EKo")
	pool.Put("Kurniawan")
	pool.Put("Khannedy")
	//memasukan 3 data ke dalam pool

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			//ambil data dari pool, ilang dari pool
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
			//memasukan data kembali ke dalam pool
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
