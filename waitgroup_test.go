package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//WaitGroup
// fitur yang digunakan untuk menunggu proses selesai dilakukan
// menandai proses goroutine, gunakan method Add(int). setelah selesai gunakan method Done().
// untuk menunggu proses selesai gunakan Wait()

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)
	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}
func TestWaitGroup(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go RunAsynchronous(&group)
	}
	group.Wait()
	fmt.Println("Selesai")
}
