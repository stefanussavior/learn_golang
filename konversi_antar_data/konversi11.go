package main

import (
	"fmt"
)

func main() {
	var byte1 = []byte{104, 97, 108, 111}
	var s = string(byte1)

	fmt.Printf("%s \n", s)
}
