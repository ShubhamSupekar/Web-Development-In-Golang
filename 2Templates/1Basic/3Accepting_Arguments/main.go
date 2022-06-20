package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang ="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`)
	nf, err := os.Create("index.html") //creating a file
	if err != nil {
		log.Fatal("error creating file", err) // checking error
	}
	defer nf.Close()                    // closing file
	io.Copy(nf, strings.NewReader(str)) //copy str into the file

}

//go run main.go Shubham
