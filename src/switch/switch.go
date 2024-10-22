package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 4:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("is the weekend !")
	default:
		fmt.Println("is a work day :(")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("no idea: %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
