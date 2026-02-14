package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

func main() {
	a := Node{value: 1} // called head actually
	b := Node{value: 2, next: &a}
	c := Node{value: 4, next: &b}
	d := Node{value: 3, next: &c}

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)

	// how to insert a new record? ✅
	// how to find a record? ✅
	// how to delete a record? ✅
	// how to update a record? ✅
	// print all records ✅
	// sort linked list

	// do I have use always the root node for all operations?
	foundNode, err := searchValueInLinkedList(&d, 8)
	if err != nil {
		fmt.Println("We couldn't find the value")
	}

	if err == nil {
		fmt.Println("We could find")
		fmt.Println(foundNode.value)
	}
	printLinkedList(&d)
	// deleteValueInLinkedList(&d, 1, nil)
	searchValueAndUpdateInLinkedList(&d, 1, 5)
	printLinkedList(&d)

}

func printLinkedList(currentNode *Node) {
	if currentNode == nil {
		return
	}
	fmt.Println(currentNode.value)
	fmt.Println(currentNode)
	printLinkedList(currentNode.next)
}

func deleteValueInLinkedList(currentNode *Node, search int, previousNode *Node) error {
	if currentNode == nil {
		return errors.New("Search do not exist in LinkedList")
	}

	if currentNode.value == search {
		fmt.Println("Found node to delete!")
		fmt.Println(currentNode)
		if previousNode == nil {
			currentNode = nil
			return nil
		}
		previousNode.next = currentNode.next
		return nil
	}
	return deleteValueInLinkedList(currentNode.next, search, currentNode)
}

func searchValueInLinkedList(currentNode *Node, search int) (*Node, error) {
	if currentNode == nil {
		return nil, errors.New("Search do not exist in LinkedList")
	}
	if currentNode.value == search {
		return currentNode, nil
	}
	return searchValueInLinkedList(currentNode.next, search)
}

func searchValueAndUpdateInLinkedList(currentNode *Node, search int, update int) error {
	if currentNode == nil {
		return errors.New("Search do not exist in LinkedList")
	}

	if currentNode.value == search {
		fmt.Println("Found node to update!")
		fmt.Println(currentNode)
		currentNode.value = update
		return nil
	}
	return searchValueAndUpdateInLinkedList(currentNode.next, search, update)
}

/*
If you want, next we can:

refactor this into a clean Go API

add insert / reverse / sort

compare this to a C implementation

or build tests for it (very Go-ish)

Just tell me 👌
*/
