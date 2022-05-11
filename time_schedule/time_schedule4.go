package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("Expired")
		ch <- true
	})

	fmt.Println("start")
	<-ch
	fmt.Println("finish")
}
