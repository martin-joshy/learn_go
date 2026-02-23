package main

import "fmt"

func main() {
	// intializing/declaring array
	var a [5]int
	b := [5]int{}
	c := [...]int{1, 2}
	d := [...]int{100, 5: 400, 500}
	fmt.Println(d)

	// 2D array
	var a2 [2][3]int

	b2 := [2][3]int{
		{1, 2},
		{3, 4, 5},
	}
}
