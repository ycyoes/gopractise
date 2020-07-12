// flowtest
package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("I am run after the function complete")
	defer fmt.Println("I am the second defer statement")
	defer fmt.Println("I am the third defer statement")

	fmt.Println("Hello World!")
}
