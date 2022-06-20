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

	Unhealthy := sugar{
		Name: "Unhealthy Diet",
		Type: "Large",
	}

	Dibetic := sugar{
		Name: "Sugar free",
		Type: "none",
	}

	SliceOfSugar := []sugar{Dibetic, Diet, Unhealthy}

	err := tpl.Execute(os.Stdout, SliceOfSugar)
	if err != nil {
		log.Fatalln(err)
	}
}
