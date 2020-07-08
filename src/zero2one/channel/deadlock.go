package main

import "fmt"

func main()  {
	pipline := make(chan string)
	pipline <- "hello world"
	fmt.Println(<-pipline)
	close(pipline)
}