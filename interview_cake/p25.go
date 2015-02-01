// K To Last Node in a Singly-Linked List
package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{
		Value: value,
		Next:  nil,
	}
}

// Returns the kth/nd last node in the list.
// Assumed n is the head of the linked list.
//
// If k = 1 the last node is returned.
//
// Assuming we don't know the length of the list.
//
// We do 2 passes through the list.
//
// Pass 1: Find the length of the list
// Pass 2: Go up to the len(l)-k Node, return it
//
func KToLastNode(n *Node, k int) *Node {
	start := n
	l := 0
	for ; n != nil; n = n.Next {
		l++
	}

	if k > l || k < 1 {
		return nil
	}

	c := 0
	for n = start; c < l-k; n = n.Next {
		c++
	}
	return n
}

var a, b, c, d, e *Node

func setup() {
	a = NewNode("A")
	b = NewNode("B")
	c = NewNode("C")
	d = NewNode("D")
	e = NewNode("E")

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
}

var tests = []struct {
	k    int
	want interface{}
}{
	{
		k:    1,
		want: "E",
	},
	{
		k:    2,
		want: "D",
	},
	{
		k:    3,
		want: "C",
	},
	{
		k:    4,
		want: "B",
	},
	{
		k:    5,
		want: "A",
	},
	{
		k:    6,
		want: nil,
	},
	{
		k:    100,
		want: nil,
	},
	{
		k:    -1,
		want: nil,
	},
}

func main() {
	setup()
	for _, t := range tests {
		got := KToLastNode(a, t.k)
		if got != t.want {
			fmt.Errorf("k %d: got %v, want %v", t.k, got, t.want)
		}
	}
}
