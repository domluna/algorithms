// Valid BST
// Dont care if it's balanced or not
package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

func Insert(n *Node, data int) *Node {
	if n == nil {
		return &Node{
			data:  data,
			left:  nil,
			right: nil,
		}
	}

	if data < n.data {
		n.left = Insert(n.left, data)
	} else {
		n.right = Insert(n.right, data)
	}
	return n
}

// We can't just check parent child relationships.
// We have to check a node in relation to the whole path.
//
// TODO: Depth first search with upper and lower bounds
//
// Path length is O(lg n), number of paths is O(2^n), so
// total runtime is O(n)
//
// O(n) space. This would only occur if our tree is
// an array so we just check one path of length n.
func CheckValid(n *Node) bool {
	return true
}

func InOrder(n *Node) {
	if n == nil {
		return
	}
	InOrder(n.left)
	fmt.Println(n.data)
	InOrder(n.right)
}

func main() {
	var root *Node
	root = Insert(root, 4)
	root = Insert(root, 3)
	root = Insert(root, 5)
	root = Insert(root, 6)
	root = Insert(root, 2)
	root = Insert(root, 1)

	//	    4
	//	   / \
	//	  3   5
	//       /     \
	//      2       6
	//     /
	//    1

	// check it's in the right order
	// 1,2,3,4,5,6
	InOrder(root)
	fmt.Println(CheckValid(root)) // true
	root.data = 15
	fmt.Println(CheckValid(root)) // false

	// reset
	root.data = 4

	root.left.data = 22
	fmt.Println(CheckValid(root)) // false
	root.left.data = 1
	fmt.Println(CheckValid(root)) // false

	// reset
	root.left.data = 3

	root.right.right.data = 1
	fmt.Println(CheckValid(root)) // false

	root.right.right.data = 10
	fmt.Println(CheckValid(root)) // true
}
