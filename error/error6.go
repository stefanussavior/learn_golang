package main

import (
	"fmt"
)

func main() {
	data := []string{"superman", "aquaman", "wonder woman"}

	for _, each := range data {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic error on looping", each, "| Message:", r)
				} else {
					fmt.Println("Application running perfectly")
				}
			}()
			panic("Some error happend")
		}()
	}
}
