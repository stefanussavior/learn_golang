package main

import (
	"fmt"
)

func main() {
	orderSomeFood("pizza")
	orderSomeFood("burger")
}

func orderSomeFood(menu string) {
	defer fmt.Println("Terima Kasih, Silahkan tunggu")
	if menu == "pizza" {
		fmt.Println("pilihan tepat!", " ")
		fmt.Println("Pizza di tempat kami paling enak", "\n")
		return
	}
	fmt.Println("pesanan anda:", menu)
}
