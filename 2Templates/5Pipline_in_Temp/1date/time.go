package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println(t)

	tf := t.Format(time.Kitchen)
	fmt.Println(tf)
	//or
	tf = t.Format("01-02-2006") //for date 01, for month 02 , for year 2006
	fmt.Println(tf)

}
