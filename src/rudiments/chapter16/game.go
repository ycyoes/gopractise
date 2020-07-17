// game
package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Guess the name of my pet to win a prize!")

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	fmt.Println("text: ", text)
	fmt.Println(text == "John")
	fmt.Println("John" == "John")
	fmt.Println(reflect.TypeOf(text))
	if text == "John" {
		fmt.Println("You won! You win chocolate!")
	} else {
		fmt.Println("You didn't win. Better luck next time")
	}
}
