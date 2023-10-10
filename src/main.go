package main

import (
	"time"

	"github.com/Ben-bo/golang-goroutines/src/channels"
)

func main() {
	channel := make(chan string)
	defer close(channel)
	go channels.ReciveData(channel)
	go channels.SendData(channel, "beni")

	time.Sleep(2 * time.Second)
}
