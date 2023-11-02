package main

import (
	"fmt"
	"log"
	"net/http"

	handlers2 "github.com/kxplxn/go-concurrency/orders-app/handlers"
)

func main() {
	fmt.Println("Welcome to the Orders App!")
	handler, err := handlers2.New()
	if err != nil {
		log.Fatal(err)
	}
	// start server
	router := handlers2.ConfigureHandler(handler)
	fmt.Println("Listening on localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
