package main

import "fmt"

func main() {
	var chiken = map[string]int{
		"Januari":  56,
		"Februari": 34,
		"Maret":    42,
		"April":    21,
		"mei":      22,
	}

	var value, isExist = chiken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item Tidak Ada")
	}
}
