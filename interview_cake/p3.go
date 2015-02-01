// Highest Product of 3
package main

import "fmt"

// a always has >= 3 elements
// Negatives!
//
// How do we get positives?
// 1. neg * pos * neg
// 2. pos * pos * pos
//
// Note: for the product of k numbers
// we need an even number of negatives in
// the product.
//
// Further is opens up for combinations
// of possible outputs.
//
// For k = 4
//
// 1. neg * 2 * pos * 2
// 2. pos * 4
// 3. neg * 4
//
// I feel a good way to get around this to ignore all the extra ones
// at first, keep track of k negatives (k-1 if k is odd) and k positives.
//
// At the end then mix and match these in all the ways and see what comes out
// on top.
//
// We could also short circuit this way. For example if k = 11:
//
// If 4 negatives * 7 positives < 2 negs * 9 * positives
// we know not to explore any further mixes.

func highestProduct(a []int) int {
	if len(a) < 3 {
		panic("len(a) < 3")
	}

	// postives
	n := []int{0, 0, 0}
	// negatives
	neg1, neg2 := 0, 0

	for _, v := range a {
		i := minIndex(n)
		if v > n[i] {
			n[i] = v
		}

		// negatives
		if v < neg1 {
			temp := neg1
			neg1 = v

			if temp < neg2 {
				neg2 = temp
			}
		}
	}

	max := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] > max {
			max = n[i]
		}
	}

	p1 := n[0] * n[1] * n[2]
	p2 := max * neg1 * neg2

	if p1 > p2 {
		return p1
	}
	return p2
}

func minIndex(a []int) int {
	i, v := 0, a[0]
	for ix := 1; ix < len(a); ix++ {
		if a[ix] < v {
			v = a[ix]
			i = ix
		}
	}
	return i
}

var tests = []struct {
	input []int
	want  int
}{
	{
		[]int{1, 2, 3},
		6,
	},
	{
		[]int{1, 0, 2, 3},
		6,
	},
	{
		[]int{-1, 0, -2, -3},
		0,
	},
	{
		[]int{1, 2, 3, -5, -10},
		150,
	},
	{
		[]int{4, 10, -4, -5, 10},
		400,
	},
	{
		[]int{1, 10, -5, 1, -100, -1},
		5000,
	},
}

func main() {
	for _, t := range tests {
		got := highestProduct(t.input)
		if got != t.want {
			fmt.Printf("ERROR input = %v: got %d, want %d\n", t.input, got, t.want)
		}
	}
}
