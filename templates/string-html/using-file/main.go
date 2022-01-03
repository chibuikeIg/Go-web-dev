package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	// create a string of html document and parse it into a file using go

	full_name := "John Doe"

	html := fmt.Sprint(`	<!DOCTYPE html>
	<html>
		<head>
		 <title>Hello world</title>
		</head>
		<body><h2>` + full_name + `</h2></body>
	</html>`)

	newHtmlFile, err := os.Create("index.html")

	if err != nil {

		fmt.Println("An Error has occured while attempting to create file", err)

	}

	defer newHtmlFile.Close()

	io.Copy(newHtmlFile, strings.NewReader(html))

}
