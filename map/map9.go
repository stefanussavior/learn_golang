package main

import "fmt"

func main() {
	var data = []map[string]string{
		{"name": "chiken blue", "gender": "male", "color": "brown"},
		{"address", "mangga street", "id": "k0001"},
		{"community", "chikens lovers"},
	}
	for _, dat := range data {
		fmt.Println(dat["name"], dat["gender"])
	}
}
