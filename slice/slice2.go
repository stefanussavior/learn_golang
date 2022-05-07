package main

import "fmt"

func main() {
	var fruits = []string{"Apple", "Grape", "Banana", "Melon"}
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]
	var aaFruits = fruits[1:2]
	var baFruits = fruits[0:1]

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)

	baFruits[0] = "pineaple"
	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)
}
