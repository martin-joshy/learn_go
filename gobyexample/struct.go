package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// create a new struct
	p1 := person{name: "martin", age: 26}
	p2 := person{"Muneera", 26}
	fmt.Println(p1, p2)

	p3 := person{name: "Jithu"}
	fmt.Println(p3)
	fmt.Println(&p3)

	dog := struct {
		breed string
		age   int
	}{"lab", 5}

	fmt.Println(dog)
}
