// Left-leaning Red-black Tree
// http://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf
// http://www.cs.princeton.edu/~rs/talks/LLRB/RedBlack.pdf
// Based on 2-3 Trees.
//
// 3-nodes are represented as a node with a left Red child.
// 3-nodes are left-leaning.
//
// 4-nodes are represented as a node with two Red children.
//
// Disallowed (invariants):
//  1. right-leaning 3-node representation
//  2. two reds in a row
//
package llrb

// Color specifies the color of a node.
type Color bool

const (
	Red   Color = false
	Black Color = true
)

// Key represents a key in in a LLRB Tree. A Key
// must be comparable
type Key interface {
	// Less returns true if a < b
	Less(a, b Key) bool

	// Equal returns true if a == b
	Equal(a, b Key) bool
}

// Node represents a node in the LLRB Tree.
type Node struct {
	Key   int
	Value interface{}
	Color

	left, right *Node
}

// Tree represents a 2-3 LLRB Tree.
type Tree struct {
	root *Node
}

// New creates a new LLRB Tree.
func New() *Tree {
	return &Tree{
		root: nil,
	}
}

// Get searches for a Node is the Tree based on a key. If the node is
// found Node.Value is returned, otherwise nil.
func (t *Tree) Get(key int) interface{} {
	return get(t.root, key)
}

func get(n *Node, key int) interface{} {
	for n != nil {
		if key == n.Key {
			return n.Value
		} else if key < n.Key {
			n = n.left
		} else {
			n = n.right
		}
	}
	return nil
}

// Height returns the height of the Tree. The height is the number
// of levels before the bottom most node is found.
func (t *Tree) Height() int {
	return height(t.root) + 1
}

func height(n *Node) int {
	if n == nil {
		return 0
	}

	leftHeight := height(n.left)
	rightHeight := height(n.right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Insert inserts a new node into the Tree based on key and value.
// This operations runs in O(log n)
func (t *Tree) Insert(key int, value interface{}) {
	t.root = insert(t.root, key, value)
	t.root.Color = Black
}

// insert inserts a node into the Tree as it would in a regular BST.
// After the insertion has been completed if an invariant has been violated
// it will be fixed, maintaining O(log n) height.
func insert(n *Node, key int, value interface{}) *Node {
	if n == nil {
		return &Node{
			Key:   key,
			Value: value,
			Color: Red,
		}
	}

	// If colors are flipped here this turns into a 2-3-4 Tree.

	if key == n.Key {
		n.Value = value
	} else if key < n.Key {
		n.left = insert(n.left, key, value)
	} else {
		n.right = insert(n.right, key, value)
	}

	n = fixUp(n)

	return n
}

// Delete deletes the node with the given key from the Tree.
func (t *Tree) Delete(key int) {
	delete(t.root, key)
}

func delete(n *Node, key int) *Node {
	if key < n.Key {
		if !isRed(n.left) && !isRed(n.left.left) {
			n = moveRedLeft(n)
			n.left = delete(n.left, key)
		}
	} else {

		if isRed(n.left) {
			// n = leanRight(n)
		}

		if key == n.Key && n.right == nil {
			return nil
		}

		if !isRed(n.right) && !isRed(n.right.left) {
			n = moveRedRight(n)
		}

		if key == n.Key {
			n.Key = min(n.right).Key
			n.Value = get(n.right, n.Key)
			n.right = deleteMin(n.right)
		} else {
			n.right = delete(n.right, key)
		}

	}

	return fixUp(n)
}

// DeleteMin deletes the minimum element of the Tree
func (t *Tree) DeleteMin() {
	t.root = deleteMin(t.root)
	t.root.Color = Black
}

func min(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return min(n.left)
}

func max(n *Node) *Node {
	if n.right == nil {
		return n
	}
	return max(n.right)
}

func deleteMin(n *Node) *Node {
	if n.left == nil {
		return nil
	}

	if !isRed(n.left) && !isRed(n.left.left) {
		n = moveRedLeft(n)
	}

	n.left = deleteMin(n.left)

	return fixUp(n)
}

// DeleteMax deletes the minimum element of the Tree
func (t *Tree) DeleteMax() {
	t.root = deleteMax(t.root)
	t.root.Color = Black
}

func deleteMax(n *Node) *Node {
	if isRed(n.left) {
		n = rotateRight(n)
	}

	if n.right == nil {
		return nil
	}

	if !isRed(n.right) && !isRed(n.right.left) {
		n = moveRedRight(n)
	}

	n.left = deleteMax(n.left)

	return fixUp(n)
}

// isRed returns true if n.Color == Red, false otherwise.
func isRed(n *Node) bool {
	if n == nil {
		return false
	}
	return n.Color == Red
}

// colorFlip flips the Color of n and its direct children.
func colorFlip(n *Node) {
	n.Color = !n.Color
	n.left.Color = !n.left.Color
	n.right.Color = !n.right.Color

}

func rotateRight(n *Node) *Node {
	x := n.left
	n.left = x.right
	x.right = n
	x.Color = n.Color
	n.Color = Red
	return x
}

func rotateLeft(n *Node) *Node {
	x := n.right
	n.right = x.left
	x.left = n
	x.Color = n.Color
	n.Color = Red
	return x
}

func moveRedLeft(n *Node) *Node {
	colorFlip(n)
	if isRed(n.right.left) {
		n.right = rotateRight(n.right)
		n = rotateLeft(n)
		colorFlip(n)
	}
	return n
}

func moveRedRight(n *Node) *Node {
	colorFlip(n)
	if isRed(n.left.left) {
		n = rotateRight(n)
		colorFlip(n)
	}
	return n
}

// fixUp fixes any violated invariants on Node n.
func fixUp(n *Node) *Node {
	// Disallowed right-leaning.
	if isRed(n.right) && !isRed(n.left) {
		n = rotateLeft(n)
	}

	// Disallowed two reds in a row.
	if isRed(n.left) && isRed(n.left.left) {
		n = rotateRight(n)
	}

	// Change if 4-node.
	if isRed(n.left) && isRed(n.right) {
		colorFlip(n)
	}
	return n
}
