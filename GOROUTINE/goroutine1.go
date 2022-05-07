package main

import (
	"fmt"
	"runtime"
)

func Print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	go Print(2, "Hallo")
	Print(4, "Apa kabar")

	var input string
	fmt.Scanln(&input)
}
