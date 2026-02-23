package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 1)

	go func() {
		c <- "Hello World"
	}()

	select {
	case msg := <-c:
		fmt.Println(msg)
	case <-time.After(time.Second):
		fmt.Println("Timeout")

	}
}
