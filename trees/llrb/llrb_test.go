package llrb_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/domluna/algorithms/trees/llrb"
)

func TestHeight(t *testing.T) {
	tree := llrb.New()

	nNodes := 1 << 20
	for i := 0; i < nNodes; i++ {
		n := rand.Int()
		tree.Insert(n, n)
	}

	h := tree.Height()
	fmt.Printf("Height of LLRB: %d, Log2: %d\n", h, 20)
}
