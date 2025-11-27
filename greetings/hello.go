package greetings

import "fmt"

// Hello return the a greetings
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
