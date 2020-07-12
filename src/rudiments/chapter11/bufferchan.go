// bufferchan
package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 2)
	messages <- "hello"
	messages <- "world"
	close(messages)
	fmt.Println("Pushed two messages onto Channel with no receivers")
	time.Sleep(time.Second * 1)
	receiver(messages)
}

func receiver(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}
