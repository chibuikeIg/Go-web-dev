package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseGlob("template/*"))

}

func main() {

	// tpl, err := template.ParseGlob("template/*")

	// if err != nil {

	// 	log.Fatalln(err)

	// }

	// err := tpl.Execute(os.Stdout, nil) // this will only execute and return the first file's content

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err := tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", nil)

	if err != nil {
		log.Fatalln(err)
	}

	// err = tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", nil)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

}
