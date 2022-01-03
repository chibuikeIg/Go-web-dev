package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var fm = template.FuncMap{

	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {

	s = strings.TrimSpace(s)
	s = s[:3]

	return s
}

var tpl *template.Template

func init() {

	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))

}

type Sage struct {
	Name  string
	Motto string
}

func main() {

	a := Sage{Name: "Buddha", Motto: "The belief of no beliefs"}

	b := Sage{Name: "Jesus", Motto: "The truth and the light"}

	slice := []Sage{a, b}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", slice)

	if err != nil {

		log.Fatalln(err)

	}

}
