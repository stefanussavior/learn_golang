package main

import (
	"fmt"
	"math"
)

func calculate(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return area, circumference
}

func main() {
	var diameter float64 = 7
	var area, circumference = calculate(diameter)
	fmt.Printf("Luas Lingkaran\t\t: %.2f \n", area)
	fmt.Printf("Keliling Lingkaran\t\t: %.2f \n", circumference)
}
