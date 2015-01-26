// Left-leaning Red-black Tree http://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf
// Based on 2-3 Trees.
// TODO: implement delete
package llrb

// Color specifies the color of a node.
type Color bool

const (
	Red   Color = false
	Black Color = true
)

// Key represents a key in in a LLRB tree. A Key
// must be comparable
type Key interface {
	// Less returns true if a < b
	Less(a, b Key) bool

	// Equal returns true if a == b
	Equal(a, b Key) bool
}

// Node represents a node in the LLRB.
type Node struct {
	Key   int
	Value interface{}
	Color

	left, right *Node
}

type Tree struct {
	root *Node
}

// New creates a new LLRB Tree.
func New() *Tree {
	return &Tree{
		root: nil,
	}
}

func (t *Tree) Search(key int) interface{} {
	n := t.root
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

func (t *Tree) Insert(key int, value interface{}) {
	t.root = insert(t.root, key, value)
	t.root.Color = Black
}

func insert(n *Node, key int, value interface{}) *Node {
	if n == nil {
		return &Node{
			Key:   key,
			Value: value,
			Color: Red,
		}
	}

	if key == n.Key {
		n.Value = value
	} else if key < n.Key {
		n.left = insert(n.left, key, value)
	} else {
		n.right = insert(n.right, key, value)
	}

	if isRed(n.right) && !isRed(n.left) {
		n = rotateLeft(n)
	}

	if isRed(n.left) && isRed(n.left.left) {
		n = rotateRight(n)
	}

	if isRed(n.left) && isRed(n.right) {
		colorFlip(n)
	}

	return n
}

func (t *Tree) DeleteMin() {
	t.root = deleteMin(t.root)
	t.root.Color = Black
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

func isRed(n *Node) bool {
	if n == nil {
		return false
	}
	return n.Color == Red
}

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

func fixUp(n *Node) *Node {
	if isRed(n.right) && !isRed(n.left) {
		n = rotateLeft(n)
	}

	if isRed(n.left) && isRed(n.left.left) {
		n = rotateRight(n)
	}

	if isRed(n.left) && isRed(n.right) {
		colorFlip(n)
	}
	return n
}
