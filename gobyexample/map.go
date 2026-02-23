package main

import "fmt"

func main() {
	zeroMap := make(map[int]int)
	zeroMap[1] = 1
	zeroMap[2] = 2

	val1 := zeroMap[1]
	fmt.Println(val1)
	fmt.Println(zeroMap)
}
