package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kxplxn/go-concurrency/01-goroutines-and-sync/03-lock/handlers"
)

func main() {
	fmt.Println("Welcome to the Orders App!")
	handler, err := handlers.New()
	if err != nil {
		log.Fatal(err)
	}
	// start server
	router := handlers.ConfigureHandler(handler)
	fmt.Println("Listening on localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
