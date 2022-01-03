package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseFiles("tpl.gohtml") // we can parse more than one files at a time

	if err != nil {
		fmt.Println(err)
	}

	newFile, err := os.Create("index.html")

	if err != nil {

		fmt.Println(err)

	}

	defer newFile.Close()

	err = tpl.Execute(newFile, nil) // When more than one file is parsed use ExecuteTemplate to specify which template to execute

	if err != nil {

		fmt.Println(err)

	}

}
