package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	// create a html document and concate data to string using data from standard output

	name := os.Args[1]

	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	html := fmt.Sprint(`	<!DOCTYPE html>
	<html>
		<head>
		 <title>Hello world</title>
		</head>
		<body><h2>` + name + `</h2></body>
	</html>`)

	htmlFile, err := os.Create("index.html")

	if err != nil {

		fmt.Println("An error has occured while creating a file", err)

	}

	defer htmlFile.Close()

	io.Copy(htmlFile, strings.NewReader(html))
}
