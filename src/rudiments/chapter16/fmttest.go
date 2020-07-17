// fmttest
package main

import (
	"fmt"
	"os"
)

type Animal struct {
	Name  string
	Color string
}

func main() {
	a := Animal{
		Name:  "Cat",
		Color: "Black",
	}
	fmt.Printf("%v\n", a)
	fmt.Printf("%+v\n", a)
	
	os.
	for i, arg := range os.Args {
		fmt.Println("argument", i, "is", arg)
	}
}
