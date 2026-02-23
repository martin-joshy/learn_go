package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "hello world, from server 1"
	}()

	go func() {
		time.Sleep(time.Second)
		c2 <- "hello world, from server 2"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("you recived msg through channel 1", msg1)
	case msg2 := <-c2:
		fmt.Println("you recived msg through channel 2", msg2)
	}
}
