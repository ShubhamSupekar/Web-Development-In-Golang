package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type person struct {
	Name string
	Age  int
}

func main() {
	p1 := person{
		Name: "Shubham",
		Age:  21,
	}

	err := tpl.Execute(os.Stdout, p1)

	if err != nil {
		log.Fatalln(err)
	}

}
