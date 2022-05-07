package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(text)

	var str = regex.ReplaceAllString(text, "potato")
	fmt.Println(str)
}
