// multiRequest
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go request(&wg)
	}
	wg.Wait()
}

func request(wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get("http://localhost:3003/xqgl/gridRelation/getWGByXQ")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)
}
