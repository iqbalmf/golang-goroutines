package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// POOL
// digunakan untuk goroutine, untuk menyimpan data selanjutnya untuk menggunakan datanya bisa menggunakan pool tersebut.
// pool dalam golang aman dari problem race condition
func RunAsync(group *sync.WaitGroup, pool *sync.Pool) {
	defer group.Done()
	group.Add(1)
	data := pool.Get()
	fmt.Println(data)
	pool.Put(data)
}
func TestPool(t *testing.T) {
	var pool = sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}
	group := sync.WaitGroup{}
	pool.Put("Iqbal")
	pool.Put("M")
	pool.Put("Fauzan")

	for i := 0; i < 10; i++ {
		go RunAsync(&group, &pool)
	}
	group.Wait()
	fmt.Println("Selesai")
}
