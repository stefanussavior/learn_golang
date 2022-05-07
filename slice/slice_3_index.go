package main

import "fmt"

func main() {
	var fruits = []string{"Apple", "Grape", "Banana"}
	var aFruis = fruits[0:2]
	var bFruits = fruits[0:2:2]

	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	fmt.Println(aFruis)
	fmt.Println(len(fruits))
	fmt.Println(cap(bFruits))

	fmt.Println(bFruits)
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))
}
