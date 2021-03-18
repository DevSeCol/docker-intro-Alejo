package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	hostName, _ := os.Hostname()
	_, err := io.WriteString(w, hostName+"<br>"+req.RemoteAddr)
	if err != nil {
		log.Fatal(err)
	}
}
