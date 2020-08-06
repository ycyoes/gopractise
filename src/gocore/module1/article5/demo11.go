package main

import (
	"fmt"
)

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero_map", 1: "one_map", 2: "two_map"}
	fmt.Printf("The element is %q.\n", container[1])
}
