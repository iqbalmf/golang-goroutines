package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// ONCE
// memastikan sebuah fungsi hanya dieksekusi sekali
// berapapun banyak goroutine yang akses, goroutine pertama yang bisa eksekusi fungsinya
// goroutine lain akan dihiraukan, tidak akan dieksekusi lagi

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
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("Counter: ", counter)
}
