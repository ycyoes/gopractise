// debug
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func main() {
	debug := os.Getenv("DEBUG")

	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		IdleConnTimeout:       90 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	/**
	* 通过创建一个传输（transport）并将其传递给客户端，可更细致地控制超时：控制HTTP连接的各个阶段。
	* 在大多数情况下，使用Timeout就足以控制整个HTTP事务，但在Go语言中，还可通过创建传输来控制HTTP事务的各个部分。
	**/
	client := &http.Client{
		// Timeout: 50 * time.Millisecond,
		Transport: tr,
	}
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
