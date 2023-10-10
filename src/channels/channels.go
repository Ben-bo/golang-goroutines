package channels

import "fmt"

// recieve data form channels
func ReciveData(ch <-chan string) {

	fmt.Println(<-ch)
}

// send data to channels
func SendData(ch chan<- string, data string) {
	ch <- data
}
