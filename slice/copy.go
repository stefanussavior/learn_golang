package main

import "fmt"

func main() {
	dst := make([]string, 3)
	src := []string{"Watermelon", "Pinnaple", "Apple", "Orange"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
}
