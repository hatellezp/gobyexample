package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")

	time1 := time.NewTimer(2 * time.Second)

	<-time1.C
	fmt.Println("Timer1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer2 stopped")
	}

	time.Sleep(2 * time.Second)
}
