// helloword
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8000", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Header.Get("Accept") {
	case "application/json":
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write([]byte("Hello World\n"))
	case "application/xml":
		w.Header().Set("Content-Type", "applicaton/xml;charset=utf-8")
		w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?><Message>Hello World</Message>`))
	default:
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.Write([]byte("Hello World\n text/plain \n"))
	}
}
