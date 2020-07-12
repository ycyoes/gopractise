// functest
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	add := addUp(2, 4)
	fmt.Printf("%v\n", add)

	fmt.Println("isEven: ", isEven(2))

	result := sumNumber(1, 2, 3, 4, 5)
	fmt.Printf("The result is %v\n", result)

	x, y := sayHi()
	fmt.Printf("say hi: %v %v\n", x, y)

	fmt.Println(feedMe(1, 0))

	fn := func() {
		fmt.Println("function called")
	}
	fn()

	anotherFunc(fn)
}

func anotherFunc(f func()) {
	f()
}

func feedMe(portion int, eaten int) int {
	eaten = portion + eaten
	if eaten >= 5 {
		fmt.Printf("I'm full! I've eaten %d\n", eaten)
		return eaten
	}
	fmt.Printf("I'm still hungry! I've eaten %d\n", eaten)
	return feedMe(portion, eaten)
}

func sayHi() (x, y string) {
	x = "hello"
	y = "world"
	return
}

func sumNumber(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//判断是否为奇数
func isEven(i int) bool {

	return i%2 == 0
}

func addUp(x int, y int) int {
	return x + y
}
