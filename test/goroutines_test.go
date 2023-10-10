package testpackage

import (
	"testing"

	"github.com/Ben-bo/golang-goroutines/src/channels"
)

func hello(name string) string {
	return name
}
func TestHello(t *testing.T) {
	go hello("test")
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	// go hello("beni")
	go channels.ReciveData(channel)
	go channels.SendData(channel, "Beni")

	// channel <- "beni"
	// dataChannel := <-channel

	// fmt.Println(dataChannel)

}
