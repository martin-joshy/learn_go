package main

import (
	"fmt"

	"example/greetings"
)

func main() {
	message := greetings.Hello("Martin")
	fmt.Println(message)
}
