package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {

	s := "Hello world this is a string which will be encoded using base64 encoding"

	encodedString := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(encodedString)

	decodedString, err := base64.StdEncoding.DecodeString(encodedString)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(decodedString))

}
