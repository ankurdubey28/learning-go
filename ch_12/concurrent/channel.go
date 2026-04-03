package main

import (
	"fmt"
	"time"
)

func main() {
	msgChan := make(chan string)
	go sendMessage(2, msgChan)
	go sendMessage(3, msgChan)

	msg1 := <-msgChan
	fmt.Println(msg1)

	msg2 := <-msgChan
	fmt.Println(msg2)

	close(msgChan)

}

func sendMessage(num int, msgChan chan<- string) {
	fmt.Printf("sending message: %d\n\n", num)
	time.Sleep(time.Second * time.Duration(num)) // simulate some work

	msg := fmt.Sprintf("Message %d! sent", num)
	msgChan <- msg
}
