package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	value int
	next  *Node
}

func add(list *Node, data int) *Node {
	if list == nil {
		list := new(Node)
		list.value = data
		list.next = nil
		return list
	}
	temp := new(Node)
	temp.value = data
	temp.next = list

	list = temp
	return list

}

func display(list *Node) {
	var temp *Node
	temp = list

	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func main() {
	var head *Node
	head = nil

	// Add 5 data points
	fmt.Println("=== Insert 5 data points ===")
	n := 0
	for n < 5 {
		fmt.Printf("Data %d\n", n)
		head = add(head, rand.Intn(100))
		n++
	}
	fmt.Println("=== Display ===")
	display(head)
}
