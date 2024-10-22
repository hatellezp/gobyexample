package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	a := func() int {
		i++
		return i
	}

	return a
}

func main() {
	fmt.Println("Hello World!")

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

}
