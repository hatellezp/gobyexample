package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			i, more := <-jobs

			if more {
				fmt.Println("received job", i)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
