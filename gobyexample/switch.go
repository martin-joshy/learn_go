package main

import "fmt"

func main() {
	// 1.classic switch
	i := 1
	switch i {
	case 1:
		fmt.Println("It is oneee")
	case 2:
		fmt.Println("It is twooo")
	default:
		fmt.Println("I don't know Mhann")
	}
	// 2.if else alternative
	switch {
	case i == 1:
		fmt.Println("You again oneee")
	default:
		fmt.Println("Whatever")
	}
	// 3.multiple experssion
	switch i {
	case 1, 0:
		fmt.Println("you againnnnnn")
	default:
		fmt.Println("Whatever")
	}
	// 4.type switch
	whoami := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("Heyyy I am a bool")
		case string:
			fmt.Println("Heyyy I am a string")
		default:
			fmt.Println("I dont know my type Mhann")
		}
	}
	whoami(1)
}
