// timeCompare
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	s1 := "2020-08-02T10:59:20+00:00"
	s2 := "2020-08-03T10:59:20+00:00"
	today, err := time.Parse(time.RFC3339, s1)
	if err != nil {
		log.Fatal(err)
	}
	tomorrow, err := time.Parse(time.RFC3339, s2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(today.After(tomorrow))
	fmt.Println(today.Before(tomorrow))
	fmt.Println(today.Equal(tomorrow))

}
