package main

import "fmt"

func main() {
	c1 := make(chan int, 5)

	for i := 0; i < 5; i++ {
		c1 <- i
	}

	close(c1)

	for elem := range c1 {
		fmt.Println(elem)
	}
}
