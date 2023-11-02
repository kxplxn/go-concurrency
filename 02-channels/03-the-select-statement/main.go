package main

import (
	"fmt"
	"time"
)

var greetings = []string{"Hello!", "Ciao!", "Hola!", "Hej!", "Salut!"}

func main() {
	// create a channel
	ch := make(chan string, 1)

	// start the greeter to provide a greeting
	go greet(ch)

	// sleep for a long time
	time.Sleep(1 * time.Second)
	fmt.Println("Main ready!")

	// receive greeting
	for greeting := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Greeting received!", greeting)
	}
}

// greet writes a greet to the given channel and then says goodbye
func greet(ch chan<- string) {
	fmt.Printf("Greeter ready!")
	for _, g := range greetings {
		ch <- g
	}
	close(ch)
	fmt.Println("Greeter completed!")
}
