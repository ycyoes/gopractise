package main

import (
	"fmt"
)

func main() {
	x := 3
	fmt.Println("x = ", x)
	x1 := add1(x)
	fmt.Println("x+1 = ", x1)
	fmt.Println("x = ", x)

	y := 3
	fmt.Println("y = ", y)
	y1 := add2(&y)
	fmt.Println("y+1 = ", y1)
	fmt.Println("y = ", y)
}

//传指针
func add2(a *int) int {
	*a = *a + 1 //修改了a的值
	return *a
}

//传值
func add1(a int) int {
	a = a + 1
	return a
}
