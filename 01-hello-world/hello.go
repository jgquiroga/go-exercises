package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(quote.Go())

	// An array
	var colors = [3]string{"Red", "Blue", "Green"}
	fmt.Println(colors)

	fmt.Println("Number of colors:", len(colors))

	// An Slice
	var colorSlice = []string{"Red", "Blue", "Green"}
	fmt.Println(colorSlice)

	// To append to a slice
	colorSlice = append(colorSlice, "Purple")
	fmt.Println(colorSlice)

	// To remove something from the slice
	colorSlice = append(colorSlice[1:len(colorSlice)])
	fmt.Println(colorSlice)

	colorSlice = append(colorSlice[:len(colorSlice)-1])
	fmt.Println(colorSlice)

	// Variable number of items
	numbers := make([]int, 5)
	numbers[0] = 1
	numbers[0] = 2
	numbers[0] = 3
	numbers[0] = 4
	numbers[0] = 5

	numbers = append(numbers, 6)

	fmt.Println(numbers)

	// Fixed number of items
	numbers = make([]int, 5, 5)
	numbers[0] = 1
	numbers[0] = 2
	numbers[0] = 3
	numbers[0] = 4
	numbers[0] = 5
	fmt.Println(numbers)

	// Dictionary (map)

	states := make(map[string]string)
	fmt.Println(states)
	states["BA"] = "Buenos Aires"
	states["CB"] = "Córdoba"
	fmt.Println(states)

}
