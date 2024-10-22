package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	fmt.Println("Hello World!")

	fmt.Println(plus(2, 4))
	fmt.Println(plusPlus(2, 4, 5))
}
