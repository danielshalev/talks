package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	// struct is initialized with 0 values
	var person1 Person

	// struct is initialized with 0 values, 
	// person2 is a pointer
	person2 := new(Person)

	person3 := Person{
		name: "John",
		age: 25,
	}

	fmt.Println(person3.name)
}