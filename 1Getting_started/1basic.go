package main

import "fmt"

type secretAgent struct {
	person
	licenseTokill bool
}

type person struct {
	first_name string
	last_name  string
}

func (sa secretAgent) speak() {
	fmt.Println(sa.first_name, sa.last_name, `says,"Shaken"`)
}

func (p person) speak() {
	fmt.Println("Hi", p.first_name, `says,"Good Morning, James."`)
}

type human interface {
	speak()
}

func saysomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		"Shubham",
		"Supekar",
	}
	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	saysomething(p1)
	saysomething(sa1)
}
