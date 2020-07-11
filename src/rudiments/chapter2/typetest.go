// type_test.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(sayHello("ycyoes"))
	var b bool
	fmt.Println(b)
	b = true
	fmt.Println(b)
}

func sayHello(s string) string {
	return "Hello, " + s
}
