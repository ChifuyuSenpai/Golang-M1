// main.go

package main

import (
	"fmt"
)

type Person struct {
	Name string
	age  int
}

func New(name string, paramAge int) Person {
	return Person{

		Name: name,
		age:  paramAge,
	}
}

func (Person) String() string {
	return "My name is " + Person.Name + Person.age + "yo"
}

func main() {
	fmt.Println(Person{})
}
