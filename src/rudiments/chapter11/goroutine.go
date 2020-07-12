// goroutine
package main

import (
	"fmt"
	"time"
)

func main() {
	slowFunc()
	fmt.Println("I am not shown until slowFunc()")
	go slowFunc()
	fmt.Println("I am now shown straightway!")
	time.Sleep(time.Second * 3)
}

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper() finished")
}
