package main

import "fmt"

func main() {
	var chikens = []map[string]string{
		{"name": "chiken blue", "gender": "male"},
		{"name": "chiken red", "gender": "male"},
		{"name": "chiken yellow", "gender": "female"},
	}

	for _, chiken := range chikens {
		fmt.Println(chiken["gender"], chiken["name"])
	}
}
