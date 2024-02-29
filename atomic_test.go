package golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

//ATOMIC
// package yang digunakan untuk menggunakan data primitive secara aman di goroutine
// contoh sebelumnya menggunakan Mutex, dalam Atomic pun bisa digunakan.
// jika hanya primitive cukup Atomic, sedangkan jika datanya struct gunakan Mutex
// banyak function di Atomic
// https://pkg.go.dev/sync/atomic

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("counter", x)
}
