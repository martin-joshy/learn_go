package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rec struct {
	length, width float64
}

type circle struct {
	radius float64
}

func (r *rec) area() float64 {
	return r.length * r.width
}

func (r *rec) perim() float64 {
	return 2 * r.length * 2 * r.width
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func getMsrmt(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r1 := rec{10, 10}
	c1 := circle{10}

	getMsrmt(&r1)
	getMsrmt(&c1)
}
