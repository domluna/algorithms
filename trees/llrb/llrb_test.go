package llrb_test

import (
	"math/rand"
	"testing"

	"github.com/domluna/algorithms/trees/llrb"
)

func TestInsertAndGet(t *testing.T) {
	tree := llrb.New()
	nNodes := 1 << 16
	keys := rand.Perm(nNodes)

	for _, k := range keys {
		tree.Insert(k, k)
	}

	for _, k := range keys {
		if v := tree.Get(k); v == nil {
			t.Errorf("Get: %v, got %v, want %v", k, nil, k)
		}
	}

}

func TestHeight(t *testing.T) {
	tree := llrb.New()
	nNodes := 1 << 16

	for i := 0; i < nNodes; i++ {
		n := rand.Int()
		tree.Insert(n, n)
	}

	h := tree.Height()
	if h >= 2*h {
		t.Errorf("Height: should be <= 2 * %d, got %d", 16, h)
	}
}

func TestDelete(t *testing.T) {
	tree := llrb.New()
	nNodes := 1 << 16
	keys := rand.Perm(nNodes)

	for _, k := range keys {
		tree.Insert(k, k)
	}

	l := tree.Len()
	if l != nNodes {
		t.Errorf("Len: got %d, want %d", l, nNodes)
	}

	// Delete
	for _, k := range keys {
		tree.Delete(k)
	}

	l = tree.Len()
	if l != 0 {
		t.Errorf("Len: got %d, want %d", l, 0)
	}

}
