// helloword
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8001", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello World\n"))
}
