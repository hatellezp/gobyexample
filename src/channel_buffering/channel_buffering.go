package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
