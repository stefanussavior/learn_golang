package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "john wick"

	var encodeString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded : ", encodeString)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encodeString)
	var decodeString = string(decodeByte)
	fmt.Println("decoded : ", decodeString)
}
