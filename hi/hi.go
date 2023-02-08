package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Hiroshi")
	fmt.Println(message)
}
