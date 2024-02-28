package golang_goroutines

import (
	"fmt"
	"testing"
)

//dapat memilih data tercepat dari beberapa channel,
//jika datang bersamaan maka akan dipilih secara random

func TestChannelSelect(t *testing.T) {
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
		}
		if counter == 2 {
			break
		}
	}

	//select {
	//case data := <-channel1:
	//	fmt.Println("Data dari channel 1", data)
	//case data := <-channel2:
	//	fmt.Println("Data dari channel 2", data)
	//
	//}
}
