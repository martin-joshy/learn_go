package main

import "fmt"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	select {
	case msg1 := <-c1:
		fmt.Println("Go the message :", msg1)
	default:
		fmt.Println("Got no message from channel 1")
	}

	select {
	case c2 <- "World":
		fmt.Println("Send World")
	default:
		fmt.Println("Could not send")
	}

	select {
	case msg1 := <-c1:
		fmt.Println("Got the message :", msg1)
	case msg2 := <-c2:
		fmt.Println("Got the message :", msg2)
	default:
		fmt.Println("Did not get any messages")
	}
}
