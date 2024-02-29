package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

//MAP
// sync.Map mirip dengan Golang MAP
// bedanya ini untuk penggunaan concurrent goroutine
// function yang bisa digunakan di MAP:
// Store(k,v) menyimpan data ke map
// Load(k) mengambil data dari map using key
// Delete(k) hapus data dari map using key
// Range(func(k,v)) untuk iterasi seluruh data map

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}
func TestMapGoroutine(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}
	group.Wait()
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ",", value)
		return true
	})
}
