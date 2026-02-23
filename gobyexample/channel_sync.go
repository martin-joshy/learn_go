package main

import (
	"fmt"
	"time"
)

func work(done chan bool) {
	fmt.Println("Starting to Work")
	time.Sleep(time.Second)
	fmt.Println("Completed the work")

	done <- true
}

func main() {
	done := make(chan bool, 1)

	go work(done)

	<-done
}
