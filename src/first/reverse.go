// reverse.go
package main

import (
	"fmt"

	. "github.com/isdamir/gotype"
)

func main() {
	fmt.Println("Hello World!")
}

func Reverse(node *LNode) {
	if node == nil || node.Next == nil return
	
}
