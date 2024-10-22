package main

import (
	"fmt"
	"iter"
	"slices"
)

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			value := yield(a)
			// fmt.Println("------------")
			// fmt.Println(a, b, value)
			if !value {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	fmt.Println("Hello World!")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(12)
	lst.Push(23)

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 40 {
			break
		}

		fmt.Println(n)
	}
}
