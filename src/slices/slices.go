package main

import (
	"fmt"
	// "slices"
)

func main() {
	fmt.Println("Hello World!")

	var s []string
	fmt.Println("unitit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("l again:", l)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i * (j + i)
		}
	}

	fmt.Println("2d:", twoD)
}
