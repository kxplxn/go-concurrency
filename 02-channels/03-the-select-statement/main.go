package main

import (
	"fmt"
	"time"
)

var hellos = []string{
	"Hello!", "Ciao!", "Hola!", "Hej!", "Salut!",
}

var goodbyes = []string{
	"Goodbye!", "Arrivederci!", "Adios!", "Hej hej!", "La revedere!",
}

func main() {
	// create a channel
	ch := make(chan string, 1)
	ch2 := make(chan string, 1)

	// start the greeter to provide a greeting
	go greet(hellos, ch)
	go greet(goodbyes, ch2)

	// sleep for a long time
	time.Sleep(1 * time.Second)
	fmt.Println("Main ready!")

	for {
		select {
		case gr, ok := <-ch:
			if !ok {
				ch = nil
				break
			}
			printGreeting(gr)
		case gr2, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			printGreeting(gr2)
		default:
			return
		}
	}
}

// greet writes a greet to the given channel and then says goodbye
func greet(messages []string, ch chan<- string) {
	fmt.Printf("Greeter ready!")
	for _, m := range messages {
		ch <- m
	}
	close(ch)
	fmt.Println("Greeter completed!")
}

func printGreeting(greeting string) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Greeting received:", greeting)
}
