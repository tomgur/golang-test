package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// request a greeting message
	message, err := greetings.Hello("Tom")
	if err != nil {
		// if no message, print error
		log.Fatal(err)
	}
	// print the message
	fmt.Println(message)
	quote := greetings.GetRandomQuote()
	fmt.Println(quote)
}
