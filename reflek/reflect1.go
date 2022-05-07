package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 20.4
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Jenis number : ", reflectValue.Type())

	if reflectValue.Kind() == reflect.Float64 {
		fmt.Println("Nilai dalam tersebut : ", reflectValue.Float())
	}
}
