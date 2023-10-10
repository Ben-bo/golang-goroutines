package test

import (
	"fmt"
	"testing"
)

func hello(name string) string {
	return name
}
func TestHello(t *testing.T) {
	go hello("test")
}

func reciveData(ch <-chan string) {

	fmt.Println(<-ch)
}
func sendData(ch chan<- string, data string) {
	ch <- data
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	// go hello("beni")
	go reciveData(channel)
	go sendData(channel, "Beni")

	// channel <- "beni"
	// dataChannel := <-channel

	// fmt.Println(dataChannel)

	close(channel)
}
