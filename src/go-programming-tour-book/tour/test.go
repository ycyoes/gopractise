// test.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
	var word = "hello"
	fmt.Println("test title: ", strings.Title(word))
	for i, r := range word {
		fmt.Println(i, string(r))
	}
}
