package main

import (
	"fmt"
)

func main() {
	var i int = 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	i = 0
	for {
		fmt.Println("loop")
		i = i + 1
		if i > 3 {
			break
		}
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}
