// type_test.go
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println(sayHello("ycyoes"))
	var b bool
	fmt.Println(b)
	b = true
	fmt.Println(b)

	var s string = "string"
	var i int = 10
	var f float32 = 1.2

	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))

	s = "true"
	b, err := strconv.ParseBool(s)
	fmt.Println(b, err)

	s = strconv.FormatBool(true)
	fmt.Println(s)

	fmt.Println(str2int("2d"))
}

func sayHello(s string) string {
	return "Hello, " + s
}

//字符串类型转int
func str2int(s string) int {
	// r, err := strconv.ParseInt(s)
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
		fmt.Println(err)
	}
	return r
}
