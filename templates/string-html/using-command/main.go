package main

import "fmt"

func main() {

	// create a string of html document and save it to a html file using commandline

	name := "John"

	html := `<!DOCTYPE html>
	<html>
	<head>
	<title>Hello world</title>
	</head>
	<body><h2>` + name + `</h2></body>
	</html>`

	fmt.Println(html)

}
