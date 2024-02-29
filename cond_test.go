package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//Cond
// implementasi locking berbasis kondisi
// cond membutuhkan Locker (Mutex / RWMutex)
// function Signal(), untuk memberitahu sebuah goroutine agar tidak perlu nunggu.
// function Broadcast(), untuk memberi tahu semua goroutine agar tidak perlu nunggu.
// untuk membuat Cond, function sync.NewCond(Locker)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	group.Add(1)
	defer group.Done()

	cond.L.Lock()
	cond.Wait()
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
			//cond.Broadcast()
		}
	}()
	group.Wait()
}
