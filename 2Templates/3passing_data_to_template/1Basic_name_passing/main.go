package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("Error creating filr", err)
	}
	defer nf.Close()
	err = tpl.Execute(nf, "Shubham")
	if err != nil {
		log.Fatalln(err)
	}
}
