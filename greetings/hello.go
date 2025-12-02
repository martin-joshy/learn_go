package greetings

import (
	"errors"
	"fmt"
)

// Hello return the a greetings
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name not provided")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
