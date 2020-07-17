// errTest
package main

import (
	// "errors"
	"log"
	"os"
)

func main() {
	// var errFatal = errors.New("We only just started and we are crashing")
	// log.Fatal(errFatal)

	f, err := os.OpenFile("example03.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	log.SetOutput(f)

	for i := 1; i <= 5; i++ {
		log.Printf("Log iteration %d", i)
	}
}
