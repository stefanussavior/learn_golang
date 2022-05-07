package main

import "fmt"

func main() {
	var chiken1 = map[string]int{"januari": 50, "februari": 40}
	var chiken2 = map[string]int{
		"januari":  50,
		"februari": 20,
	}

	fmt.Println(chiken1)
	fmt.Println(chiken2)
}
