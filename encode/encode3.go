package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "https://kalipare.com/"

	var encodeString = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodeString)

	var decodeByte, _ = base64.URLEncoding.DecodeString(encodeString)
	var decodedString = string(decodeByte)
	fmt.Println(decodedString)
}
