// string.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	editStr()
	edit2()
}

//字符串切片
func edit2() {
	s := "hello"
	s = "c" + s[1:]
	fmt.Printf("%s\n", s)
}

//Go语言中字符串是不可变的，可按如下方式修改
func editStr() {
	s := "hello"
	c := []byte(s) //将字符串s转换为[]byte类型
	c[0] = 'c'
	s2 := string(c) //再转换回string类型
	fmt.Printf("%s\n", s2)
}
