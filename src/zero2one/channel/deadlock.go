package main

import "fmt"

func main()  {
	pipline := make(chan string)
	go hello(pipline)
	
	pipline <- "hello world"
	ele := <- pipline
	fmt.Printf("ele data %v\n",  ele)
	close(pipline)
}

func hello(pipline chan string)  {
	<-pipline
}