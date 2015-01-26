package llrb_test

import (
	"math/rand"
	"testing"

	"github.com/domluna/algorithms/trees/llrb"
)

func TestHeight(t *testing.T) {
	tree := llrb.New()

	nNodes := 1 << 16
	for i := 0; i < nNodes; i++ {
		n := rand.Int()
		tree.Insert(n, n)
	}

	h := tree.Height()
	if h >= 2*h {
		t.Errorf("LLRB height should be <= 2 * %d, got %d", 16, h)
	}
}

func TestSearch(t *testing.T) {
	tree := llrb.New()
	nNodes := 1 << 16
	keys := rand.Perm(nNodes)

	for _, k := range keys {
		tree.Insert(k, k)
	}

	for _, k := range keys {
		if _, ok := tree.Search(k); !ok {
			t.Errorf("Search %v, got %v, want %v", k, ok, true)
		}
	}
}

func TestDelete(t *testing.T) {

}
