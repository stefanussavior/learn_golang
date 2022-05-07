package main

import "fmt"

func main() {
	var chikens = []map[string]string{
		map[string]string{"name": "chiken blue", "gender": "male"},
		map[string]string{"name": "chiken red", "gender": "male"},
		map[string]string{"name": "chiken yellow", "gender": "female"},
	}

	for _, chiken := range chikens {
		fmt.Println(chiken["gender"], chiken["name"])
	}
}
