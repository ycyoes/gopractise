// vartest
package main

import (
	"fmt"
)

const greeting string = "Hello, world!"

func main() {
	var i int
	var f float32
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	if s == "" {
		fmt.Printf("s has not been assigned a value and is zero valued.\n")
	}

	s = "Hello world"
	fmt.Printf("Print 's' variable from outer block %v\n", s)
	b = true
	if b {
		fmt.Printf("Printing 'b' variable from outer block %v\n", b)
		i := 42
		if b != false {
			fmt.Printf("Printing 'i' variable from outer block %v\n", i)
		}
	}

	fmt.Printf("%v\n", &s)
	fmt.Println("greeting: ", greeting)
}
