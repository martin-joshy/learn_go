package main

import "fmt"

func main() {
	// 3 types of go for loop
	// 1. while loop equalant
	i := 4
	for i > 1 {
		fmt.Println(i)
		i--
	}
	// 2. classic for loop
	for i := 1; i < 7; i++ {
		fmt.Println(i)
	}
	// 3. range
	for i := range 4 {
		fmt.Println(i)
	}
}
