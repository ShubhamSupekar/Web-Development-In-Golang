package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sugar struct {
	Name string
	Type string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {

	Diet := sugar{
		Name: "For diet",
		Type: "No sugar",
	}

	err := tpl.Execute(os.Stdout, Diet)
	if err != nil {
		log.Fatalln(err)
	}
}
