package main

import (
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Why did the programmer quit his job?\n Because he didn't get arrays.")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
