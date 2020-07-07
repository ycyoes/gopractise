package main

import "fmt"

func main()  {
	
	ch1 := make(chan int, 3)  // 声明并初始化了一个元素类型为 int、容量为 3 的通道 ch1
	go func()  {
		ch1 <- 2
		ch1 <- 1
		ch1 <- 3
		elem1 := <-ch1
		fmt.Printf("The first element received from channel ch1: %v\n", elem1)
		close(ch1)
	}()

	for{
		ele, ok := <- ch1 // ok 为 false，说明通道已关闭
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		fmt.Printf("receive element %v\n", ele)
	}
	fmt.Println("end.")
}

