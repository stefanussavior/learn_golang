package main

import (
	"fmt"
)

func main() {
	orderSomeFood("pizza")
	orderSomeFood("burger")
}

func orderSomeFood(menu string) {
	defer fmt.Println("Terima kasih, silahkan tunggu")
	if menu == "burger" {
		fmt.Println("Pilihan Tepat", " ")
		fmt.Println("Pizza ditempat kami paling enak", "\n")
		return
	}
	fmt.Println("pesanan anda:", menu)
}
