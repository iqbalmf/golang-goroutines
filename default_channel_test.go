package golang_goroutines

import (
	"fmt"
	"testing"
)

//Default Channel
//sama seperti select channel, hanya penambahan default

func TestDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("menunggu Data")
		}
		if counter == 2 {
			break
		}
	}
}
