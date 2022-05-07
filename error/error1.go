package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Println("Type some number: ")
	fmt.Scan(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}
}
