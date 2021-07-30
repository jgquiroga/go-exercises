package main

import (
	"fmt"
)

func main() {

	poodle := Dog{"Poodle", 10}

	fmt.Println(poodle)

	fmt.Printf("%+v\n", poodle)

	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	value := "The current value"
	tree := Node{Value: value, Left: nil, Right: nil}

	valueLeft := "The left value"
	left := Node{Value: valueLeft, Left: nil, Right: nil}
	left.Value = valueLeft
	tree.Left = &left

	fmt.Println(tree)

	fmt.Printf("Left node: %v\n", tree.Left.Value)
}

// Dog is a struct
type Dog struct {
	// Uppercase, exported
	Breed  string
	Weight int
}

type Node struct {
	Value string
	Left  *Node
	Right *Node
}
