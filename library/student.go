package library

import "fmt"

var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "Stefanus Andre"
	Student.Grade = 2

	fmt.Println("--> library/library.go imported")
}
