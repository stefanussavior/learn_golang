package main

import "fmt"

var allStudents = []struct {
	person
	grade int
}{
	{person: person{"wick", 21}, grade: 2},
	{person: person{"ethan", 22}, grade: 3},
	{person: person{"bond", 21}, grade: 3},
}

func main() {
	for _, student := range allStudents {
		fmt.Println(student)
	}
}
