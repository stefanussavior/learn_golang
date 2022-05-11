package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
	}

	var allStudent = []person{
		{name: "Stefanus", age: 19},
		{name: "Andre", age: 20},
		{name: "Marta", age: 21},
	}

	for _, student := range allStudent {
		fmt.Println(student.name, "age is", student.age)
	}
}
