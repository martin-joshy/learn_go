package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case tick := <-ticker.C:
				fmt.Println("Tick at :", tick)
			case <-done:
				fmt.Println("Completed ticking")
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Stoped the ticker")
}
