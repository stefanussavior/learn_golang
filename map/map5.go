package main

import (
	"fmt"
)

func main() {
	var chiken = map[string]int{
		"Januari":  50,
		"Februari": 20,
	}

	fmt.Println(len(chiken))
	fmt.Println("\t", chiken)

	delete(chiken, "Januari")

	fmt.Println(len(chiken))
	fmt.Println("\t", chiken)
}
