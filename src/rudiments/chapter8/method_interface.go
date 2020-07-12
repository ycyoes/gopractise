// method_interface
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	m := Movie{
		Name:   "Spiderman",
		Rating: 3.2,
	}
	fmt.Println(m.summary())

	s := Sphere{
		Radius: 5,
	}
	fmt.Println(s.SurfaceArea())
	fmt.Println(s.Volume())

	t := Triangle{base: 3, height: 1}
	fmt.Println(t.area())

	tr := Triangle{base: 3, height: 1}
	tr.changeBase(4)
	fmt.Println(tr.base)
}

func (t Triangle) changeBase(f float64) {
	t.base = f
	return
}

func (t *Triangle) area() float64 {
	return 0.5 * (t.base * t.height)
}

func (s *Sphere) Volume() float64 {
	radiusCubed := s.Radius * s.Radius * s.Radius
	return (float64(4) / float64(3)) * math.Pi * radiusCubed
}

//表面积
func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}

func (m *Movie) summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + ", " + r
}

type Triangle struct {
	base   float64
	height float64
}

type Sphere struct {
	Radius float64
}

type Movie struct {
	Name   string
	Rating float64
}
