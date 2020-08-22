// main.go
package main

import (
	"log"

	"github.com/go-programming-tour-book/tour/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("cmd.Execute err: %v", err)
	}
}
