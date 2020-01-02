// fir
package main

import (
	"fmt"
)

type iAdder func(int) (int iAdder)

func main() {
	fmt.Println("Hello World!")

	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + 2 + ... + %d = %d\n", i, a(i))
	}
}

//返回函数
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
