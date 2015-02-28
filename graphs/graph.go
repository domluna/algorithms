// Note meant to be run
package main

// Adjacency List
type Graph1 struct {
	nodes []*Node
}

type Data struct {
	// some stuff here
}

type Node struct {
	id   int
	data Data

	// out links
	links []*Node
}

type Edge struct {
	in, out *Node
}

// Edge List O(E) space
type Graph2 struct {
	edges []*Edge
}
