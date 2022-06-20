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
func main() {

	sugar := []string{"Medium", "small", "Powder", "Rock"}
	err := tpl.Execute(os.Stdout, sugar)
	if err != nil {
		log.Fatalln(err)
	}
}
