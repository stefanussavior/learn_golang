package main

import (
	"fmt"
	"time"
)

func main() {
	var timer = time.NewTimer(5 * time.Second)
	fmt.Println("start")
	<-timer.C
	fmt.Println("Finish")
}
