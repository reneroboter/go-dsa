package main

import "fmt"

type BinaryTree struct {
	left  *BinaryTree
	right *BinaryTree
	value int
}

func main() {
	root := BinaryTree{value: 1}
	addValueToBinaryTree(&root, 2)
	addValueToBinaryTree(&root, 3)
	addValueToBinaryTree(&root, 4)
	addValueToBinaryTree(&root, 5)
	addValueToBinaryTree(&root, 6)
	addValueToBinaryTree(&root, 7)
	addValueToBinaryTree(&root, 8)
	addValueToBinaryTree(&root, 9)
	addValueToBinaryTree(&root, 10)

	printBinaryTree(root)
}

// how to find a node?
// how to delete a node?
// how to update a node?

func printBinaryTree(node BinaryTree) {
	fmt.Println(node.value)
	if node.left != nil {
		fmt.Println("left")
		printBinaryTree(*node.left)
	}

	if node.right != nil {
		fmt.Println("right")
		printBinaryTree(*node.right)
	}
}

func addValueToBinaryTree(node *BinaryTree, value int) {
	if node.right != nil && node.left != nil {
		addValueToBinaryTree(node.left, value)
		return
	}
	newNode := BinaryTree{value: value}
	if node.left != nil {
		node.right = &newNode
		return
	}
	node.left = &newNode
}
