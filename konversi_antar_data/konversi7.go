package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "false"
	var bul, err = strconv.ParseBool(str)

	if err == nil {
		fmt.Println(bul)
	}
}
