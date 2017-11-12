package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person) CanDrinkBeer() bool {
	return p.age >= 18
}

func main() {
	aPerson := Person{
		name: "John",
		age: 25,
	}

	fmt.Println(aPerson.CanDrinkBeer())
}