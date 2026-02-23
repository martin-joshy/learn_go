package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	fmt.Println("length of the slice of bytes: ", (len(s)))
	fmt.Println("length of the runes: ", utf8.RuneCountInString(s))

	// Hex values
	fmt.Println("Hexadecimal values")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v : %v : %x \n", i, s[i], s[i])
	}
	fmt.Printf("% x", s)

	// Runes
	for idx, runeValue := range s {
		fmt.Printf("%v : %q \n", idx, runeValue)
	}
	for i, w := 0, 0; i < len(s); i += w {
		runeValues, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("unicode %+q of width %v \n", runeValues, width)
		w = width
	}
}
