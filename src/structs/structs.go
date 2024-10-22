package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 43

	return &p
}

func main() {
	fmt.Println("Hello World!")

	fmt.Println(person{"bob", 20})
	fmt.Println(person{name: "alice", age: 39})
	fmt.Println(person{name: "fred"})
	fmt.Println(&person{name: "ann", age: 30})
	fmt.Println(newPerson("john"))

	s := person{name: "sean", age: 39}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"rex",
		true,
	}
	fmt.Println(dog)
}
