package main

import (
	"fmt"
)

func sum(nums ...int) int {
	fmt.Print(nums, " ")

	var total int

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println(sum(3, 6, 1, 5))

	numbers := []int{2, 5, 6}
	fmt.Println((sum(numbers...)))
}
