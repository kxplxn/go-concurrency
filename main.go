package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", Hello)
	fmt.Println("listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("failed to serve on port 8080")
	}
}

func Hello(w http.ResponseWriter, _ *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello world!"); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
