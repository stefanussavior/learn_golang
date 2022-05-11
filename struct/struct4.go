package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	age   int
	grade int
}

func main() {
	var s1 = student{}
	s1.name = "wick"
	s1.age = 21
	s1.person.age = 22

	fmt.Println(s1.name)
	fmt.Println(s1.age)
	fmt.Println(s1.person.age)
}
