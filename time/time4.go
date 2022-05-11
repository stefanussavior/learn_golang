package main

import (
	"fmt"
	"time"
)

func main() {
	var date, err = time.Parse("06 Jan 15", "02 sep 15 08:00")

	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	fmt.Println(date)
}
