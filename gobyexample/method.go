package main

import "fmt"

type rect struct {
	length, width int
}

func (r *rect) area() int {
	return r.length * r.width
}

func (r *rect) modify() (int, error) {
	r.length = 10
	r.width = 20
	return 1, nil
}

func main() {
	r1 := rect{5, 2}

	fmt.Println(r1.area())
	fmt.Println(r1)
	fmt.Println(r1.modify())
	fmt.Println(r1)
}
