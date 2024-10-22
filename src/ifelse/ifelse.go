package main

import (
	"fmt"
)

func main() {

	var seven int = 7
	if seven%2 == 0 {
		fmt.Println(seven, "is even")
	} else {
		fmt.Println(seven, "is odd")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
