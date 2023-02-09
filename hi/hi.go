package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Takeshi", "Satoshi", "Ivankov"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(messages)
	for name, message := range messages {
		fmt.Printf("Hey, %v!%v\n", name, message)
	}
}
