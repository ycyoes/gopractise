// importtest
package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/golang/example/stringutil"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(math.Pi)
	fmt.Println(strings.ToLower("STOP SHOUTING"))
	s := "ti esrever dna ti pilf nwod gniht ym tup I"
	fmt.Println(stringutil.Reverse(s))
}
