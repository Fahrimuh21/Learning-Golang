package belajar_golang_goroutines

//tentang penggunaan sync.Map pada goroutine
//dokumentasi sync.Map: https://pkg.go.dev/sync#Map

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	//untuk menambahkan data ke dalam sync.Map
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}
	//membuat sync.Map dan sync.WaitGroup

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
		//menjalankan 100 goroutine untuk menambahkan data ke dalam sync.Map
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
