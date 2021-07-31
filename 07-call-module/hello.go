package main

import (
	"fmt"

	"jgqsolutions.com.ar/greetings"
)

func main() {
	message, err := greetings.Hello("Julian")
	fmt.Println(message)

	message, err = greetings.Hello("")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(message)
	}

}
