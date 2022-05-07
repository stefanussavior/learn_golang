package main

import (
	"fmt"
	"math"
)

func calculate(d float64) (float64, float64) {
	//hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	//hitung keliling
	var circumate = math.Pi * d

	//kembalikan 2 nilai
	return area, circumate
}

func main() {
	var diameter float64 = 15
	var area, circumate = calculate(diameter)

	fmt.Printf("Luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("Keliling Lingkaran\t: %.2f \n", circumate)
}
