package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	nums := []int{2, 3, 4}
	var sum int

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "gogo" {
		fmt.Println(i, c)
	}
}
