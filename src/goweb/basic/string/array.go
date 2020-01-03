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

	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Printf("sum is equal to %d", sum)

	//for可简写为：
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println("sum2: ", sum2)

	//进一步简化
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println("sum3 ", sum3)

}
