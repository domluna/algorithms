// Goal: decompose the tree to make the biggest forest.
// Catch is every tree in the forest has to have an
// even number of vertices.
//
// Observations:
//
// A tree with an odd # of vertices can't be broken
// down to two trees of even vertices.
//
// odd = even + odd | odd + even
//
// So a tree with odd # of vertices can't be broken
// down any further.
//
// Idea:
//
// Find the maximum even cut in a tree, recurse on
// the two trees.
//
// Why would this work?
//
// Not sure!
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Since we have at most ~10,000 edges
// we could represent this as a 2D array.
// Tree
type Tree struct {
	links map[int][]int
	cache map[int]int
}

func newTree() *Tree {
	return &Tree{
		links: make(map[int][]int),
		cache: make(map[int]int),
	}
}

// evenTree takes away the maximum number of edges
// if subtree subTreeSize is even we can cut that edge
func evenTree(t *Tree) int {
	// find the root of the whole tree
	// the root has no incoming edges
	s := 0
	for _, l := range t.links {
		for _, c := range l {
			if t.treeSize(c)%2 == 0 {
				s += 1
			}
		}
	}

	return s
}

func (t *Tree) addEdge(u, v int) {
	if _, ok := t.links[u]; ok {
		t.links[u] = append(t.links[u], v)
		return
	}

	t.links[u] = make([]int, 0)
	t.links[u] = append(t.links[u], v)
}

func (t *Tree) treeSize(root int) int {
	// see if cached
	if v, ok := t.cache[root]; ok {
		return v
	}

	children := t.links[root]
	s := 1
	for _, c := range children {
		s += t.treeSize(c)
	}

	// set size before returning
	t.cache[root] = s
	return s
}

func main() {

	var N, M int
	tree := newTree()

	fmt.Scan(&N) // 2 <= N <= 100
	fmt.Scan(&M) // number of edges

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < M; i++ {
		scanner.Scan()
		u, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		tree.addEdge(v, u) // nodes are read in a weird way
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// works
	// fmt.Println(tree.treeSize(1))
	// fmt.Println(tree.treeSize(2))
	// fmt.Println(tree.treeSize(3))
	// fmt.Println(tree.treeSize(6))
	// fmt.Println(tree.treeSize(10))

	fmt.Println(evenTree(tree))
}
