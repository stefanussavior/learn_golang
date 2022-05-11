package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	time.Sleep(5 * time.Second)

	duration := time.Since(start)

	fmt.Println("time elapsed in second : ", duration.Seconds())
	fmt.Println("time elapsed in minute : ", duration.Minutes())
	fmt.Println("time elapsed in hour : ", duration.Hours())
}
