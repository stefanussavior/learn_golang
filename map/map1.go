package main

import "fmt"

func main() {
	var chiken map[string]int
	chiken = map[string]int{}

	chiken["januari"] = 50
	chiken["februari"] = 20

	fmt.Println("Januari", chiken["januari"])
	fmt.Println("Februari", chiken["februari"])
}
