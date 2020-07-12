// struct_point
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	m := Movie{
		Name:   "Citizen Kane",
		Rating: 10,
	}
	fmt.Println(m.Name, m.Rating)
	fmt.Printf("%+v\n", m)

	m.Name = "Metropolis"
	m.Rating = 0.9918
	fmt.Printf("%+v\n", m)

	e := Superhero{
		Name: "Batman",
		Age:  32,
		Address: Address{
			Number: 1007,
			Street: "Mountain Driver",
			City:   "Gotham",
		},
	}
	fmt.Printf("%+v\n", e)

	fmt.Printf("%+v\n", NewAlarm("07:00"))

	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}

	b := Drink{
		Name: "Lemonade",
		Ice:  true,
	}

	if a == b {
		fmt.Println("a and b are the same")
	}
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)

	c := a
	c.Ice = false
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", c)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)

	//要修改原始结构体实例包含的值，必须使用指针。
	// 指针是指向内存地址的引用，因此使用它操作的不是结构体的副本而是其本身。
	// 要获得指针，可在变量名前加上和号
	d := &a
	d.Ice = false
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", d)
	fmt.Printf("%+v\n", *d)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", d)
}

type Drink struct {
	Name string
	Ice  bool
}

func NewAlarm(time string) Alarm {
	a := Alarm{
		Time:  time,
		Sound: "Klaxon",
	}
	return a
}

type Alarm struct {
	Time  string
	Sound string
}

type Address struct {
	Number int
	Street string
	City   string
}

type Superhero struct {
	Name    string
	Age     int
	Address Address
}

type Movie struct {
	Name   string
	Rating float32
}
