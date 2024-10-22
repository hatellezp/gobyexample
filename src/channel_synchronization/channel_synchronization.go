package main

import (
	"fmt"
	"time"
)

func some_function() {
	time.Sleep(time.Second)
}

func worker(done chan bool, f func()) {
	fmt.Print("working ...")
	f()
	fmt.Println("done")

	done <- true

}

func main() {
	fmt.Println("Hello World!")

	done := make(chan bool, 1)
	go worker(done, some_function)

	<-done
}
