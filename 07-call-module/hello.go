package main

import (
	"fmt"

	"jgqsolutions.com.ar/greetings"
)

func main() {
	message := greetings.Hello("Julian")
	fmt.Println(message)
}
