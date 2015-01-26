// Implements traversal algorithms for binary trees.
// the binary tree implements focuses on ints for
// simplicity.
//
// A more complete example might use an interface defined
// by a compare method.
//
package main

import "fmt"

// tree represents a binary tree.
type tree struct {
	left, right *tree
	value       int
}

// insert inserts a value into tree and returns a modified
// version of the tree.
func insert(t *tree, value int) *tree {
	if t == nil {
		return &tree{
			value: value,
		}
	}

	if value > t.value {
		t.right = insert(t.right, value)
		return t
	}
	t.left = insert(t.left, value)
	return t
}

// traverseAlg represents a function for traversing a tree.
type traverseAlg func(t *tree) []int

// inOrder traverses the tree in the left, middle, right
// order.
func inOrder(t *tree) []int {
	if t == nil {
		return []int{}
	}
	nodes := []int{}
	left := inOrder(t.left)
	right := inOrder(t.right)

	nodes = append(nodes, left...)
	nodes = append(nodes, t.value)
	nodes = append(nodes, right...)
	return nodes
}

// postOrder traverses the tree in left, right, middle
// order.
func postOrder(t *tree) []int {
	if t == nil {
		return []int{}
	}
	nodes := []int{}
	left := postOrder(t.left)
	right := postOrder(t.right)

	nodes = append(nodes, left...)
	nodes = append(nodes, right...)
	nodes = append(nodes, t.value)
	return nodes
}

// preOrder traverses the tree in middle, left, right
// order.
func preOrder(t *tree) []int {
	if t == nil {
		return []int{}
	}
	nodes := []int{}
	left := preOrder(t.left)
	right := preOrder(t.right)

	nodes = append(nodes, t.value)
	nodes = append(nodes, left...)
	nodes = append(nodes, right...)
	return nodes
}

// traverse traverses the tree based on passed traverseAlg.
func traverse(t *tree, alg traverseAlg) {
	values := alg(t)
	fmt.Println("Values:", values)
}

func main() {
	var t *tree

	t = insert(t, 6)
	t = insert(t, 7)
	t = insert(t, 2)
	t = insert(t, 1)
	t = insert(t, 4)
	t = insert(t, 3)
	t = insert(t, 5)
	t = insert(t, 9)
	t = insert(t, 8)

	fmt.Println("preorder")
	traverse(t, preOrder)
	fmt.Println("inorder")
	traverse(t, inOrder)
	fmt.Println("postorder")
	traverse(t, postOrder)
}
