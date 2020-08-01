package main

import (
	"io/ioutil"
	"log"
)

func main() {
	s := "hello world"
	content := []byte(s)
	err := ioutil.WriteFile("example03.txt", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
