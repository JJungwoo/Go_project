package main

import (
	"fmt"
)

type Person struct {
	name	string
	age		int
}

func (p Person) Hello() {
	fmt.Printf("Hello. My name is %s\n", p.name)
}

func (p Person) MyAge() {
	fmt.Printf("My age is %d\n", p.age)
}

func main() {
	human := Person{name: "jungwoo", age: 29}
	human.Hello()
	human.MyAge()
}


