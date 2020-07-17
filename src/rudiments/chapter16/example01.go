// example01
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println("argument", i, "is", arg)
	}

}
