package main

import (
	"fmt"
)

func example1() *Node {
	node6 := Node{Value: 6, Left: nil, Right: nil}
	node5 := Node{Value: 5, Left: nil, Right: &node6}
	node4 := Node{Value: 4, Left: nil, Right: nil}
	node3 := Node{Value: 3, Left: nil, Right: nil}
	node2 := Node{Value: 2, Left: &node3, Right: &node4}
	node1 := Node{Value: 1, Left: &node2, Right: &node5}

	binaryTree := node1
	return &binaryTree
}

func main() {

	tree := example1()

	fmt.Println("* right branch without flatten:")
	tree.printAsLinkedList()

	tree.flatten()

	fmt.Println("* right branch flatten:")
	tree.printAsLinkedList()

}

type Node struct {
	Value       int32
	Left, Right *Node
}

func (node *Node) printAsLinkedList() {

	fmt.Println(node.Value)
	if node.Right != nil {
		node.Right.printAsLinkedList()
	}

}

//  Exercise taken from:
//  Flatten a binary tree into linked list
//  https://www.geeksforgeeks.org/flatten-a-binary-tree-into-linked-list/

//  Explanation:
//  The flatten method is used recursively to create a linked list in place.
//  Returns the last node so we can use it in a high order Node

func (node *Node) flatten() *Node {
	left := node.Left
	right := node.Right

	lastNode := node

	if left != nil {
		lastNode = left.flatten()

		// Last Node replaces the right node
		lastNode.Right = right
		node.Right = left
	}

	// Reset left
	node.Left = nil

	if right != nil {
		lastNode = right.flatten()
	}

	return lastNode

}
