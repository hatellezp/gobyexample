package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c3 <- "three"
	}()

	// note that the third message is not waited upon as
	// the `for` loop only has two loops
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case msg3 := <-c3:
			fmt.Println("received", msg3)
		}
	}
}
