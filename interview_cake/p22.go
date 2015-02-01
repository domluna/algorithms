// Delete Node
package main

import "fmt"

// Node represents a node in a linked list.
type Node struct {
	Value interface{}
	Next  *Node
}

// deleteNode deletes a node from a linked list.
// Note: the last node can't truly be deleted,
// all we can do is set its value to nil.
func deleteNode(n *Node) *Node {
	// empty node
	if n == nil {
		return nil
	}

	nn := n.Next
	// tail node
	if nn == nil {
		n.Value = nil
		return n
	}

	n.Value = nn.Value
	n.Next = nn.Next
	return n
}

// forward return the next node in the list.
// If there is no next node, nil is returned.
func (n *Node) forward() *Node {
	if n == nil {
		return nil
	}

	return n.Next
}

var a, b, c *Node

func setup() {
	a = &Node{
		Value: 1,
	}
	b = &Node{
		Value: 2,
	}
	c = &Node{
		Value: 3,
	}
	a.Next = b
	b.Next = c
}

func iter() {
	for n := a; n != nil; n = n.forward() {
		fmt.Printf(" %v ", n.Value)
	}
	fmt.Println()
}

func main() {
	setup()
	iter() // 1,2,3

	setup()
	a = deleteNode(a)
	iter() // 2,3

	setup()
	b = deleteNode(b)
	iter() // 1,3

	setup()
	c = deleteNode(c)
	iter() // 1,2,(nil)

}
