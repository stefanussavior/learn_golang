package main

import (
	"fmt"
	"os"
)

func main() {
	var agrsRaw = os.Args
	fmt.Printf("-> %#v\n", agrsRaw)

	var agrs = agrsRaw[1:]
	fmt.Printf("-> %#v\n", agrs)
}
