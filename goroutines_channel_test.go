package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Iqbal Fauzan"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Iqbal Fauzan"

}
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Iqbal Fauzan"
}
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestChannelInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go OnlyIn(channel)
	go OnlyOut(channel)
	time.Sleep(5 * time.Second)
}

func TestCheckPointer(t *testing.T) {
	x := 5
	xn := Square(&x)
	xs := tSquare(x)

	y := x
	fmt.Println(x, xn, xs)
	fmt.Println("y", y, "pointer y", modifyPointer(&y), "value y", modifyValue(y))
}
func tSquare(num int) int {
	return num * num
}
func Square(num *int) int {
	return *num * *num
}
func modifyPointer(num *int) int {
	*num = 10
	return *num
}
func modifyValue(num int) int {
	num = 15
	return num
}
