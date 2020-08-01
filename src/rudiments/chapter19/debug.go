// debug
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	debug := os.Getenv("DEBUG")
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://ifconfig.co", nil)
	//设置返回信息的格式
	request.Header.Add("Accept", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	if debug == "1" {
		debugRequest, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("request - %s", debugRequest)
	}

	response, err := client.Do(request)
	defer response.Body.Close()

	if debug == "1" {
		debugResponse, err := httputil.DumpResponse(response, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("response - %s", debugResponse)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("body - %s\n", body)
}
