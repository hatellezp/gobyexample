package main

import (
	"fmt"
)

func vals() (int, int) {
	return 3, 6
}

func main() {
	fmt.Println("Hello World!")

	a, b := vals()
	fmt.Println(a, b)
}
