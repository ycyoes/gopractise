// stringtest
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "I am an interpreted string literal"
	fmt.Println(s)

	s1 := `After a backslash, certain single character escapes
	represent
		special values
		 n is a line feed or new line
			t is a tab`
	fmt.Println(s1)

	s2 := "Oh sweet ignition" + " be my fuse"
	fmt.Println(s2)

	s3 := "Can you hear me?"
	s3 += "\nHear me scremain?"
	fmt.Println(s3)

	var buffer bytes.Buffer
	for i := 0; i < 500; i++ {
		buffer.WriteString("z")
	}
	fmt.Println(buffer.String())
	fmt.Println(len(buffer.String()))

	fmt.Println(strings.ToLower("VERY IMPORTANT MESSAGE"))

	//查找子串
	fmt.Println(strings.Index("surface", "face"))
	fmt.Println(strings.Index("moon", "aer"))

	fmt.Println(strings.TrimSpace(" I don't need all this space "))
}
