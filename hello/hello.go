package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("[greetings.go]: ")
	log.SetFlags(3)

	message, err := greetings.Hello("sds")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
