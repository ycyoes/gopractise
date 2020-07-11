// type_test.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(sayHello("ycyoes"))
}

func sayHello(s string) string {
	return "Hello, " + s
}
