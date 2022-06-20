package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type person struct {
	Name string
	Age  int
}

type tokill struct {
	person
	Kill bool
}

func main() {

	p1 := tokill{
		person{
			Name: "Savitor",
			Age:  66,
		},
		false,
	}
	err := tpl.Execute(os.Stdout, p1)

	if err != nil {
		log.Fatalln(err)
	}
}
