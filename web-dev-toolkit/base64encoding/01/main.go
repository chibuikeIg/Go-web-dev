package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	encoding := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	s := "Hello world this is a string which will be encoded using base64 encoding"

	encodedString := base64.NewEncoding(encoding).EncodeToString([]byte(s))

	fmt.Println(encodedString)
}
