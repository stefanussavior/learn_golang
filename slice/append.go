package main

import "fmt"

func main() {
	var fruits = []string{"Apple", "Grape", "Banana"}
	var bFruits = fruits[0:2]
	fmt.Println(cap(bFruits))
	fmt.Println(len(bFruits))

	var cFruits = append(bFruits, "papaya")
	fmt.Println(fruits)
	fmt.Println(bFruits)
	fmt.Println(cFruits)
}
