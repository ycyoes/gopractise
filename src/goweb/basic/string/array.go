package main

import (
	"fmt"
)

func main() {
	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	a := [3]int{1, 2, 3}
	fmt.Println(a)
	b := [...]int{4, 5, 6} //省略长度
	fmt.Println(b)
	fmt.Printf("The first element is %d\n", arr[0])
	fmt.Printf("The last element is %d\n", arr[9])
}
