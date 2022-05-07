package main

import (
	"fmt"
)

var numberA int = 4
var numberB *int = &numberA

func main() {
	fmt.Println("NumberA (Value) : ", numberA)
	fmt.Println("NumberA (address) : ", &numberA)

	fmt.Println("Number B (value) : ", *numberB)
	fmt.Println("Number B (address): ", numberB)
}
