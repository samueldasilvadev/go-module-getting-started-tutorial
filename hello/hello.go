package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func init() {
	log.SetPrefix("[hello.go]: ")
	log.SetFlags(3)
}

func main() {
	message, err := greetings.Hello("sds")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Test...")
	fmt.Println(message)
}
