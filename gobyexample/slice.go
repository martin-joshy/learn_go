package main

import "fmt"

func main() {
	var s []string
	fmt.Println("Nil Slice", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("empty", s, "len", len(s), "cap", cap(s))

	s[0] = "a"
	s[1] = "b"

	s = append(s, "c", "d")
	fmt.Println("len: ", len(s), "cap: ", cap(s))

	c := make([]string, len(s))
	copy(c, s)

	l := s[2:5]
	fmt.Println(l)

	t := []string{"a", "b", "c"}
	fmt.Println(t)

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
