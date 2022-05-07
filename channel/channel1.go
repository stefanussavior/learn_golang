package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var message = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		message <- data
	}

	go sayHelloTo("Stefanus Andre")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var message1 = <-message
	fmt.Println(message1)

	var message2 = <-message
	fmt.Println(message2)

	var message3 = <-message
	fmt.Println(message3)
}
