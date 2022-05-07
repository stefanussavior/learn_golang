package main

import (
	"fmt"
	"strings"
)

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ",")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My Hobbies are : %s\n", hobbiesAsString)
}

func main() {
	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)
}
