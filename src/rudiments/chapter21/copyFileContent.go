// copyFileContent
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	from, err := os.Open("./example03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile("./example03.copy.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}

}
