package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//Timer
// representasi kejadian, ketika timer expire maka event dikirim ke channel
// gunakan time.NewTimer(duration)
// contoh jika ingin mengirim kejadian tetapi nanti diwaktu tertentu bisa gunakan Timer

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())
	time := <-timer.C
	fmt.Println(time)
}

// time.After()
// Timer yang digunakan ketika tidak butuh timer hanya butuh channelnya saja
func TestTimeAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())
	time := <-channel
	fmt.Println(time)
}

// time.AfterFunc()
// kadang ada kebutuhan untuk menjalankan function dengan delay waktu tertentu
// tidak perlu lagi menggunakan channel nya, cukup kirim func yang akan dipanggil ketika timer kirim kejadian

func TestTimerAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Execute after 5 second")
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())
	group.Wait()
}

// time.Ticker()
// ticker representasi kejadian berulang
// ketika ticker expire, maka event dikirim ke channel
// untuk membuat ticker, time.NewTicker(duration)
// untuk menghentikan ticker, Ticker.Stop()

func TestTimeTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	/*Deadlock*/
	for time2 := range ticker.C {
		fmt.Println(time2)
	}

}

// time.Tick()
// kadang tidak butuh Ticker nya, butuh hanya channel saja
// untuk penggunaanya, function timer.Tick(duration)
// func ini tidak mengembalikan Ticker, hanya mengembalikan channel timer saja

func TestTimeTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)
	for time2 := range channel {
		fmt.Println(time2)
	}
}
