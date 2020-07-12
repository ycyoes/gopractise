// slice_reflect.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	var cheeses [2]string
	cheeses[0] = "Mariolles"
	cheeses[1] = "Epoisses de Bourgogen"
	fmt.Println(cheeses[0])
	fmt.Println(cheeses[1])
	fmt.Println(cheeses)

	cheeses1 := make([]string, 2)
	cheeses1[0] = "Mariolles"
	cheeses1[1] = "Epoisses de Bourgogen"
	cheeses1 = append(cheeses1, "Camembert")
	fmt.Println(cheeses1[2])

	cheeses1 = append(cheeses1, "Reblochon", "Picodon")

	fmt.Println(len(cheeses1))
	cheeses1 = append(cheeses1[:2], cheeses1[2+1:]...)
	fmt.Println(len(cheeses1))
	fmt.Println(cheeses1)

	var smellyCheeses = make([]string, 2)
	copy(smellyCheeses, cheeses1)
	fmt.Println(smellyCheeses)

	var players = make(map[string]int)
	players["cook"] = 32
	players["bairstow"] = 27
	players["stokes"] = 26
	fmt.Println(players["cook"])
	fmt.Println(players["bairstow"])
	delete(players, "cook")
	fmt.Println(players["cook"])
}
