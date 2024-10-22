package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
