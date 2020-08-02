package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println(time.Now())

	time.Sleep(3 * time.Second)
	fmt.Println("I'm awake")

	s := "2006-02-02T15:04:05+07:00"
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

	fmt.Printf("The hour is %v\n", t.Hour())
	fmt.Printf("The minute is %v\n", t.Minute())
	fmt.Printf("The second is %v\n", t.Second())
	fmt.Printf("The day is %v\n", t.Day())
	fmt.Printf("The month is %v\n", t.Month())
	fmt.Printf("The unix time is %v\n", t.Unix())
	fmt.Printf("The day of the week is %v\n", t.Weekday())

	nt := t.Add(2 * time.Second)
	fmt.Println(nt)

	formatTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("format: ", formatTime)
}
