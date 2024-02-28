package golang_goroutines

import (
	"fmt"
	"testing"
)

//Buffered Channel
// - digunakan untuk menampung data antrian di Channel,
// - karena default channel hanya bisa menerima 1 data.
// - jika menambah data ke-2, akan diminta menunggu sampai data ke-1 ada yang ambil
// - contoh: ada dimana pengirim lebih cepat dari penerima, jika menggunakan channel, maka pengirim akan melambat prosesnya

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2) //INSERT CAPACITY
	defer close(channel)
	channel <- "iqbal"
	channel <- "fauzen"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println("Selesai")
	fmt.Println(cap(channel))
	fmt.Println(len(channel))
}
