package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	s := "Hello world this is a string which will be encoded using base64 encoding"

	encodedString := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(encodedString)

}
