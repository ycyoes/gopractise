// fir
package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type iAdder func(int) (int, iAdder)

func main() {
	fmt.Println("Hello World!")

	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + 2 + ... + %d = %d\n", i, a(i))
	}

	fmt.Println("---函数式编程---")

	b := adderV2(0)
	fmt.Print(b)
	for i := 0; i < 10; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("0 + 1 + 2 + ... + %d = %d\n", i, s)
	}

	fmt.Println("---为函数实现接口---")
	f := Fibonacci()
	printFileContents(f)
}

//为函数实现接口
type Generate func() int

func Fibonacci() Generate {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func (gen Generate) Read(p []byte) (n int, err error) {
	nextNum := gen()
	numString := fmt.Sprintf("%d \n", nextNum)
	return strings.NewReader(numString).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func adderV2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adderV2(base + v)
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
