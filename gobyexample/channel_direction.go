package main

import "fmt"

func pings(ping chan<- string, msg string) {
	ping <- msg
}

func pongs(pong chan<- string, ping <-chan string) {
	msg := <-ping
	pong <- msg
}

func main() {
	ping := make(chan string, 1)
	pong := make(chan string, 1)

	pings(ping, "Hello World")
	pongs(pong, ping)

	fmt.Println(<-pong)
}
