package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectType.NumField(); i++ {
		fmt.Println("Nama : ", reflectType.Field(i).Name)
		fmt.Println("Tipe data : ", reflectType.Field(i).Type)
		fmt.Println("nilai : ", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func main() {
	var s1 = &student{Name: "Stefanus", Grade: 2}
	s1.getPropertyInfo()
}
