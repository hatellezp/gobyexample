package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("Hello World!")

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 33

	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	n2 := map[string]int{"foo": 1, "bar": 2}

	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
