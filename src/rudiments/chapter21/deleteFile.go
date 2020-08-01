// deleteFile
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//获取当前路径
	dir, _ := os.Getwd()
	fmt.Println(dir)
	err := os.Remove(dir + "/deletename.txt")
	if err != nil {
		log.Fatal(err)
	}
}
